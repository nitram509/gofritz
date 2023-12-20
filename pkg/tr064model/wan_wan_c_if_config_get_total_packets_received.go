package tr064model

import "encoding/xml"

// GetTotalPacketsReceivedResponse AUTO-GENERATED (do not edit) model from [wancommonifconfigSCPD],
// based on SOAP action 'GetTotalPacketsReceived', Fritz!Box-System-Version 164.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
type GetTotalPacketsReceivedResponse struct {
	XMLName              xml.Name // rather for debug purpose
	TotalPacketsReceived int      `xml:"NewTotalPacketsReceived"` // default=0
}
