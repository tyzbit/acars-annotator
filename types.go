package main

const (
	FlightAwareRoot        = "https://flightaware.com/live/flight/"
	PlaneSpottersPhotoRoot = "https://www.planespotters.net/photos/reg/"
)

type ACARSAnnotator interface {
	AnnotateACARSMessage(ACARSMessage) ACARSAnnotation
	Name() string
}

type Receiver interface {
	SubmitACARSMessage(AnnotatedACARSMessage) error
	Name() string
}

type ACARSAnnotation struct {
	Annotator string
	Annotation
}

type Annotation map[string]any

type AnnotatedACARSMessage struct {
	ACARSMessage
	Annotations []ACARSAnnotation
}

// This is the format ACARSHub sends
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
		ACARSRouterUUID    string `json:"acars_router_UUID"`
	}
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
