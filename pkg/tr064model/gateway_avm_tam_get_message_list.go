package tr064model

import "encoding/xml"

// GetMessageListResponse AUTO-GENERATED (do not edit) model from [x_tamSCPD],
// based on SOAP action 'GetMessageList', Fritz!Box-System-Version 164.08.00
//
// [x_tamSCPD]: http://fritz.box:49000/x_tamSCPD.xml
type GetMessageListResponse struct {
	XMLName xml.Name // rather for debug purpose
	URL     string   `xml:"NewURL"` //
}
