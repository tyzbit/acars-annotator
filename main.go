package main

import (
	"strings"

	cfg "github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
	log "github.com/sirupsen/logrus"
)

var (
	config            = Config{}
	enabledAnnotators = []ACARSAnnotator{}
	enabledReceivers  = []Receiver{}
	enabledFilters    = []string{}
)

// Set via ENV variables or a .env file
type Config struct {
	ACARSHubHost                     string `env:"ACARSHUB_HOST"`
	ACARSHubPort                     int    `env:"ACARSHUB_PORT"`
	AnnotateACARS                    bool   `env:"ANNOTATE_ACARS"`
	ADSBExchangeAPIKey               string `env:"ADBSEXCHANGE_APIKEY"`
	ADSBExchangeReferenceGeolocation string `env:"ADBSEXCHANGE_REFERENCE_GEOLOCATION"`
	FilterCriteriaMatchTailCode      string `env:"FILTER_CRITERIA_MATCH_TAIL_CODE"`
	FilterCriteriaHasText            bool   `env:"FILTER_CRITERIA_HAS_TEXT"`
	LogLevel                         string `env:"LOGLEVEL"`
	NewRelicLicenseKey               string `env:"NEW_RELIC_LICENSE_KEY"`
	NewRelicLicenseCustomEventType   string `env:"NEW_RELIC_CUSTOM_EVENT_TYPE"`
	WebhookURL                       string `env:"WEBHOOK_URL"`
	WebhookMethod                    string `env:"WEBHOOK_METHOD"`
	WebhookHeaders                   string `env:"WEBHOOK_HEADERS"`
	DiscordWebhookURL                string `env:"DISCORD_WEBHOOK_URL"`
}

// Set up Config, logging
func init() {
	// Read from .env and override from the local environment
	dotEnvFeeder := feeder.DotEnv{Path: ".env"}
	envFeeder := feeder.Env{}

	_ = cfg.New().AddFeeder(dotEnvFeeder).AddStruct(&config).Feed()
	_ = cfg.New().AddFeeder(envFeeder).AddStruct(&config).Feed()

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	loglevel := strings.ToLower(config.LogLevel)
	switch loglevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	// Add annotators based on what's enabled
	if config.AnnotateACARS {
		enabledAnnotators = append(enabledAnnotators, ACARSHandlerAnnotator{})
	}
	if config.ADSBExchangeAPIKey != "" {
		log.Info("ADSB handler enabled")
		if config.ADSBExchangeAPIKey == "" {
			log.Error("ADSB API key not set")
		}
		enabledAnnotators = append(enabledAnnotators, ADSBHandlerAnnotator{})
	}
	if len(enabledAnnotators) == 0 {
		log.Warn("no annotators are enabled")
	}

	// Add receivers based on what's enabled
	if config.WebhookURL != "" {
		log.Info("Webhook receiver enabled")
		enabledReceivers = append(enabledReceivers, WebhookHandlerReciever{})
	}
	if config.NewRelicLicenseKey != "" {
		log.Info("New Relic reciever enabled")
		enabledReceivers = append(enabledReceivers, NewRelicHandlerReciever{})
	}
	if config.DiscordWebhookURL != "" {
		log.Info("New Relic reciever enabled")
		enabledReceivers = append(enabledReceivers, DiscordHandlerReciever{})
	}
	if len(enabledReceivers) == 0 {
		log.Warn("no receivers are enabled")
	}

	// Add filters based on what's enabled
	if config.FilterCriteriaMatchTailCode != "" {
		enabledFilters = append(enabledFilters, "MatchesTailCode")
	}

	if config.FilterCriteriaHasText {
		enabledFilters = append(enabledFilters, "HasText")
	}

	SubscribeToACARSHub()
}
