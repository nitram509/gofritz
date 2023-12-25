package tr064model

import "encoding/xml"

// GetTotalBytesSentResponse AUTO-GENERATED (do not edit) model from [wancommonifconfigSCPD],
// based on SOAP action 'GetTotalBytesSent', Fritz!Box-System-Version 141.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
type GetTotalBytesSentResponse struct {
	XMLName        xml.Name // rather for debug purpose
	TotalBytesSent int      `xml:"NewTotalBytesSent"` // default=0
}
