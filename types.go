package main

const (
	FlightAwareRoot   = "https://flightaware.com/live/flight/"
	FlightAwarePhotos = "https://www.flightaware.com/photos/aircraft/"
	WebhookUserAgent  = "github.com/tyzbit/acars-annotator"
)

// Set via ENV variables or a .env file
type Config struct {
	ACARSHubHost                     string  `env:"ACARSHUB_HOST"`
	ACARSHubPort                     int     `env:"ACARSHUB_PORT"`
	AnnotateACARS                    bool    `env:"ANNOTATE_ACARS"`
	ACARSHubVDLM2Host                string  `env:"ACARSHUB_HOST"`
	ACARSHubVDLM2Port                int     `env:"ACARSHUB_VDLM2_PORT"`
	AnnotateVDLM2                    bool    `env:"ANNOTATE_VDLM2"`
	TAR1090URL                       string  `env:"TAR1090_URL"`
	TAR1090ReferenceGeolocation      string  `env:"TAR1090_REFERENCE_GEOLOCATION"`
	ACARSAnnotatorSelectedFields     string  `env:"ACARS_ANNOTATOR_SELECTED_FIELDS"`
	ADSBExchangeAPIKey               string  `env:"ADBSEXCHANGE_APIKEY"`
	ADSBExchangeReferenceGeolocation string  `env:"ADBSEXCHANGE_REFERENCE_GEOLOCATION"`
	ADSBAnnotatorSelectedFields      string  `env:"ADSB_ANNOTATOR_SELECTED_FIELDS"`
	VDLM2AnnotatorSelectedFields     string  `env:"VDLM2_ANNOTATOR_SELECTED_FIELDS"`
	TAR1090AnnotatorSelectedFields   string  `env:"TAR1090_ANNOTATOR_SELECTED_FIELDS"`
	FilterCriteriaHasText            bool    `env:"FILTER_CRITERIA_HAS_TEXT"`
	FilterCriteriaMatchTailCode      string  `env:"FILTER_CRITERIA_MATCH_TAIL_CODE"`
	FilterCriteriaMatchFlightNumber  string  `env:"FILTER_CRITERIA_MATCH_FLIGHT_NUMBER"`
	FilterCriteriaMatchFrequency     float64 `env:"FILTER_CRITERIA_MATCH_FREQUENCY"`
	FilterCriteriaMatchASSStatus     string  `env:"FILTER_CRITERIA_MATCH_ASSSTATUS"`
	FilterCriteriaAboveSignaldBm     float64 `env:"FILTER_CRITERIA_ABOVE_SIGNAL_DBM"`
	FilterCriteriaBelowSignaldBm     float64 `env:"FILTER_CRITERIA_BELOW_SIGNAL_DBM"`
	FilterCriteriaMatchStationID     string  `env:"FILTER_CRITERIA_MATCH_STATION_ID"`
	FilterCriteriaMore               bool    `env:"FILTER_CRITERIA_MORE"`
	FilterCriteriaAboveDistanceNm    float64 `env:"FILTER_CRITERIA_ABOVE_DISTANCE_NM"`
	FilterCriteriaBelowDistanceNm    float64 `env:"FILTER_CRITERIA_Below_DISTANCE_NM"`
	FilterCriteriaEmergency          bool    `env:"FILTER_CRITERIA_EMERGENCY"`
	LogLevel                         string  `env:"LOGLEVEL"`
	NewRelicLicenseKey               string  `env:"NEW_RELIC_LICENSE_KEY"`
	NewRelicLicenseCustomEventType   string  `env:"NEW_RELIC_CUSTOM_EVENT_TYPE"`
	WebhookURL                       string  `env:"WEBHOOK_URL"`
	WebhookMethod                    string  `env:"WEBHOOK_METHOD"`
	WebhookHeaders                   string  `env:"WEBHOOK_HEADERS"`
	DiscordWebhookURL                string  `env:"DISCORD_WEBHOOK_URL"`
}

type ACARSAnnotator interface {
	Name() string
	AnnotateACARSMessage(ACARSMessage) Annotation
	SelectFields(Annotation) Annotation
}

type VDLM2Annotator interface {
	Name() string
	AnnotateVDLM2Message(VDLM2Message) Annotation
	SelectFields(Annotation) Annotation
}

// ALL KEYS MUST BE UNIQUE AMONG ALL ANNOTATORS
type Annotation map[string]interface{}

type Receiver interface {
	SubmitACARSAnnotations(Annotation) error
	Name() string
}

type ACARSFilter interface {
	Filter(ACARSMessage) bool
}

// This is the format ACARSHub sends for ACARS messages
type ACARSMessage struct {
	FrequencyMHz float64 `json:"freq"`
	Channel      int     `json:"channel"`
	ErrorCode    int     `json:"error"`
	SignaldBm    float64 `json:"level"`
	Timestamp    float64 `json:"timestamp"`
	App          struct {
		Name               string `json:"name"`
		Version            string `json:"version"`
		Proxied            bool   `json:"proxied"`
		ProxiedBy          string `json:"proxied_by"`
		ACARSRouterVersion string `json:"acars_router_version"`
		ACARSRouterUUID    string `json:"acars_router_uuid"`
	} `json:"app"`
	StationID        string `json:"station_id"`
	ASSStatus        string `json:"assstat"`
	Mode             string `json:"mode"`
	Label            string `json:"label"`
	BlockID          string `json:"block_id"`
	Acknowledge      any    `json:"ack"` // Can be bool or string
	AircraftTailCode string `json:"tail"`
	MessageText      string `json:"text"`
	MessageNumber    string `json:"msgno"`
	FlightNumber     string `json:"flight"`
}

// This is the format ACARSHub sends
type VDLM2Message struct {
	VDL2 struct {
		App struct {
			Name               string `json:"name"`
			Version            string `json:"ver"`
			Proxied            bool   `json:"proxied"`
			ProxiedBy          string `json:"proxied_by"`
			ACARSRouterVersion string `json:"acars_router_version"`
			ACARSRouterUUID    string `json:"acars_router_uuid"`
		} `json:"app"`
		AVLC struct {
			CR          string `json:"cr"`
			Destination struct {
				Address string `json:"addr"`
				Type    string `json:"type"`
			} `json:"dst"`
			FrameType string `json:"frame_type"`
			Source    struct {
				Address string `json:"addr"`
				Type    string `json:"type"`
				Status  string `json:"status"`
			} `json:"src"`
			RSequence int  `json:"rseq"`
			SSequence int  `json:"sseq"`
			Poll      bool `json:"poll"`
			ACARS     struct {
				Error                 bool   `json:"err"`
				CRCOK                 bool   `json:"crc_ok"`
				More                  bool   `json:"more"`
				Registration          string `json:"reg"`
				Mode                  string `json:"mode"`
				Label                 string `json:"label"`
				BlockID               string `json:"blk_id"`
				Acknowledge           any    `json:"ack"`
				FlightNumber          string `json:"flight"`
				MessageNumber         string `json:"msg_num"`
				MessageNumberSequence string `json:"msg_num_seq"`
				MessageText           string `json:"msg_text"`
			} `json:"acars"`
		} `json:"avlc"`
		BurstLengthOctets    int     `json:"burst_len_octets"`
		FrequencyHz          int     `json:"freq"`
		Index                int     `json:"idx"`
		FrequencySkew        float64 `json:"freq_skew"`
		HDRBitsFixed         int     `json:"hdr_bits_fixed"`
		NoiseLevel           float64 `json:"noise_level"`
		OctetsCorrectedByFEC int     `json:"octets_corrected_by_fec"`
		SignalLevel          float64 `json:"sig_level"`
		Station              string  `json:"station"`
		Timestamp            struct {
			UnixTimestamp int `json:"sec"`
			Microseconds  int `json:"usec"`
		} `json:"t"`
	} `json:"vdl2"`
}
