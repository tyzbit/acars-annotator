package main

// Interface function to satisfy ACARSHandler
// Although this is the ACARS annotator, we must support ACARS and VLM2
// message types
func (a ACARSHandlerAnnotator) AnnotateVDLM2Message(m VDLM2Message) (annotation Annotation) {
	annotation = Annotation{
		"vdlm2AppName":               m.VDL2.App.Name,
		"vdlm2AppVersion":            m.VDL2.App.Version,
		"vdlm2AppProxied":            m.VDL2.App.Proxied,
		"vdlm2AppProxiedBy":          m.VDL2.App.ProxiedBy,
		"vdlm2AppRouterVersion":      m.VDL2.App.ACARSRouterVersion,
		"vdlm2AppRouterUUID":         m.VDL2.App.ACARSRouterUUID,
		"vdlmCR":                     m.VDL2.AVLC.CR,
		"vdlmDestinationAddress":     m.VDL2.AVLC.Destination.Address,
		"vdlmDestinationType":        m.VDL2.AVLC.Destination.Type,
		"vdlmFrameType":              m.VDL2.AVLC.FrameType,
		"vdlmSourceAddress":          m.VDL2.AVLC.Source.Address,
		"vdlmSourceType":             m.VDL2.AVLC.Source.Type,
		"vdlmSourceStatus":           m.VDL2.AVLC.Source.Status,
		"vdlmRSequence":              m.VDL2.AVLC.RSequence,
		"vdlmSSequence":              m.VDL2.AVLC.SSequence,
		"vdlmPoll":                   m.VDL2.AVLC.Poll,
		"vdlm2BurstLengthOctets":     m.VDL2.BurstLengthOctets,
		"vdlm2FrequencyHz":           m.VDL2.FrequencyHz,
		"vdlm2Index":                 m.VDL2.Index,
		"vdlm2FrequencySkew":         m.VDL2.FrequencySkew,
		"vdlm2HDRBitsFixed":          m.VDL2.HDRBitsFixed,
		"vdlm2NoiseLevel":            m.VDL2.NoiseLevel,
		"vdlm2OctetsCorrectedByFEC":  m.VDL2.OctetsCorrectedByFEC,
		"vdlm2SignalLevel":           m.VDL2.SignalLevel,
		"vdlm2Station":               m.VDL2.Station,
		"vdlm2Timestamp":             m.VDL2.Timestamp.UnixTimestamp,
		"vdlm2TimestampMicroseconds": m.VDL2.Timestamp.Microseconds,
		// These fields are identical to ACARS, so they will have the ACARS prefix
		"acarsErrorCode":             m.VDL2.AVLC.ACARS.Error,
		"acarsCRCOK":                 m.VDL2.AVLC.ACARS.CRCOK,
		"acarsMore":                  m.VDL2.AVLC.ACARS.More,
		"acarsAircraftTailCode":      m.VDL2.AVLC.ACARS.Registration,
		"acarsMode":                  m.VDL2.AVLC.ACARS.Mode,
		"acarsLabel":                 m.VDL2.AVLC.ACARS.Mode,
		"acarsBlockID":               m.VDL2.AVLC.ACARS.BlockID,
		"acarsAcknowledge":           m.VDL2.AVLC.ACARS.Acknowledge,
		"acarsFlightNumber":          m.VDL2.AVLC.ACARS.FlightNumber,
		"acarsMessageNumber":         m.VDL2.AVLC.ACARS.MessageNumber,
		"acarsMessageNumberSequence": m.VDL2.AVLC.ACARS.MessageNumberSequence,
		"acarsMessageText":           m.VDL2.AVLC.ACARS.MessageText,
	}
	return annotation
}
