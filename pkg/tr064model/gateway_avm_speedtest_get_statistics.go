package tr064model

// GetAvmSpeedtestStatisticsResponse AUTO-GENERATED (do not edit) model from [x_speedtestSCPD],
// based on SOAP action 'GetStatistics', Fritz!Box-System-Version 164.07.57
//
// [x_speedtestSCPD]: http://fritz.box:49000/x_speedtestSCPD.xml
type GetAvmSpeedtestStatisticsResponse struct {
	ByteCount    int `xml:"NewByteCount"`    // default=0
	KbitsCurrent int `xml:"NewKbitsCurrent"` // default=0
	KbitsAvg     int `xml:"NewKbitsAvg"`     // default=0
	PacketCount  int `xml:"NewPacketCount"`  // default=0
	PPSCurrent   int `xml:"NewPPSCurrent"`   // default=0
	PPSAvg       int `xml:"NewPPSAvg"`       // default=0
}
