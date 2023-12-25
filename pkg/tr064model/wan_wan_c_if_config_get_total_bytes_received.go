package tr064model

import "encoding/xml"

// GetTotalBytesReceivedResponse AUTO-GENERATED (do not edit) model from [wancommonifconfigSCPD],
// based on SOAP action 'GetTotalBytesReceived', Fritz!Box-System-Version 141.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
type GetTotalBytesReceivedResponse struct {
	XMLName            xml.Name // rather for debug purpose
	TotalBytesReceived int      `xml:"NewTotalBytesReceived"` // default=0
}
