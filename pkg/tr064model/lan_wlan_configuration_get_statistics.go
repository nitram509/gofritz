package tr064model

import "encoding/xml"

// GetWlanConfigurationStatisticsResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetStatistics', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetWlanConfigurationStatisticsResponse struct {
	XMLName              xml.Name // rather for debug purpose
	TotalPacketsSent     int      `xml:"NewTotalPacketsSent"`     // default=0
	TotalPacketsReceived int      `xml:"NewTotalPacketsReceived"` // default=0
}
