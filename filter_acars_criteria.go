package main

import (
	"regexp"
)

type ACARSCriteriaFilter struct {
}

func (a ACARSCriteriaFilter) Name() string {
	return "CriteriaFilter"
}

// All filters are defined here
var (
	filterFunctions = map[string]func(m ACARSMessage) bool{
		"HasText": func(m ACARSMessage) bool {
			re := regexp.MustCompile("[\\S]+")
			return re.MatchString(m.MessageText)
		},
		"MatchesTailCode": func(m ACARSMessage) bool {
			return config.FilterCriteriaMatchTailCode == m.AircraftTailCode
		},
		"MatchesFlightNumber": func(m ACARSMessage) bool {
			return config.FilterCriteriaMatchFlightNumber == m.FlightNumber
		},
		"MatchesFrequency": func(m ACARSMessage) bool {
			return config.FilterCriteriaMatchFrequency == m.FrequencyMHz
		},
		"MatchesStationID": func(m ACARSMessage) bool {
			return config.FilterCriteriaMatchStationID == m.StationID
		},
		"AboveMinimumSignal": func(m ACARSMessage) bool {
			return config.FilterCriteriaAboveSignaldBm < m.SignaldBm
		},
	}
)

// Return true if a message passes a filter, false otherwise
func (f ACARSCriteriaFilter) Filter(m ACARSMessage) (ok bool, failedFilters []string) {
	ok = true
	for _, filter := range enabledFilters {
		match := filterFunctions[filter](m)
		// Every filter must pass
		if config.FilterCriteriaInclusive && !match {
			failedFilters = append(failedFilters, filter)
			return false, failedFilters
		}
		// Any filter can pass
		if !config.FilterCriteriaInclusive && match {
			return true, failedFilters
		}
		if match == false {
			ok = false
			failedFilters = append(failedFilters, filter)
		}
	}
	return ok, failedFilters
}
