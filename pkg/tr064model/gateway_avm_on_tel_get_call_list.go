package tr064model

import "encoding/xml"

// GetCallListResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'GetCallList', Fritz!Box-System-Version 164.08.00
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetCallListResponse struct {
	XMLName     xml.Name // rather for debug purpose
	CallListURL string   `xml:"NewCallListURL"` //
}
