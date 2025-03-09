package main

import (
	"bufio"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const dictionary = "https://github.com/dwyl/english-words/raw/refs/heads/master/words.txt"

func ConfigureFilters() {
	// -------------------------------------------------------------------------

	resp, err := http.Get(dictionary)
	if err != nil {
		log.Errorf("error fetching dictionary: %v", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		englishDictionary = append(englishDictionary, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading dictionary: %v", err)
	}

	log.Debugf("loaded %d words\n", len(englishDictionary))
	// Add filters based on what's enabled
	if config.FilterCriteriaMatchTailCode != "" {
		enabledFilters = append(enabledFilters, "MatchesTailCode")
	}
	if config.FilterCriteriaHasText {
		enabledFilters = append(enabledFilters, "HasText")
	}
	if config.FilterCriteriaMatchFlightNumber != "" {
		enabledFilters = append(enabledFilters, "MatchesFlightNumber")
	}
	if config.FilterCriteriaMatchFrequency != 0.0 {
		enabledFilters = append(enabledFilters, "MatchesFrequency")
	}
	if config.FilterCriteriaMatchStationID != "" {
		enabledFilters = append(enabledFilters, "MatchesStationID")
	}
	if config.FilterCriteriaAboveSignaldBm != 0.0 {
		enabledFilters = append(enabledFilters, "AboveMinimumSignal")
	}
	if config.FilterCriteriaBelowSignaldBm != 0.0 {
		enabledFilters = append(enabledFilters, "BelowMaximumSignal")
	}
	if config.FilterCriteriaAboveDistanceNm != 0.0 {
		enabledFilters = append(enabledFilters, "AboveMinimumSignal")
	}
	if config.FilterCriteriaBelowDistanceNm != 0.0 {
		enabledFilters = append(enabledFilters, "BelowMaximumSignal")
	}
	if config.FilterCriteriaMatchASSStatus != "" {
		enabledFilters = append(enabledFilters, "MatchesASSStatus")
	}
	if config.FilterCriteriaMore {
		enabledFilters = append(enabledFilters, "More")
	}
	if config.FilterCriteriaEmergency {
		enabledFilters = append(enabledFilters, "Emergency")
	}
	if float64(config.FilterCriteriaEnglishWordCountGreaterThan) != 0 {
		enabledFilters = append(enabledFilters, "DictionaryWordCount")
	}
}
