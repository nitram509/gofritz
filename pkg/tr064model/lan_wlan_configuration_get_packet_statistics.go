package tr064model

// GetPacketStatisticsResponse auto generated model from [wlanconfigSCPD],
// based on SOAP action 'GetPacketStatistics', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetPacketStatisticsResponse struct {
	TotalPacketsSent     int `xml:"NewTotalPacketsSent"`     // default=0
	TotalPacketsReceived int `xml:"NewTotalPacketsReceived"` // default=0
}
