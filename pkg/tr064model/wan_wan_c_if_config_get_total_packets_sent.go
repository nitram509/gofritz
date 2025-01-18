package tr064model

import "encoding/xml"

// GetTotalPacketsSentResponse AUTO-GENERATED (do not edit) model from [wancommonifconfigSCPD],
// based on SOAP action 'GetTotalPacketsSent', Fritz!Box-System-Version 164.08.00
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
type GetTotalPacketsSentResponse struct {
	XMLName          xml.Name // rather for debug purpose
	TotalPacketsSent int      `xml:"NewTotalPacketsSent"` // default=0
}
