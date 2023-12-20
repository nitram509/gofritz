package tr064model

// GetWlanConfigurationStatisticsResponse auto generated model from [wlanconfigSCPD],
// based on SOAP action 'GetStatistics', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetWlanConfigurationStatisticsResponse struct {
	TotalPacketsSent     int `xml:"NewTotalPacketsSent"`     // default=0
	TotalPacketsReceived int `xml:"NewTotalPacketsReceived"` // default=0
}