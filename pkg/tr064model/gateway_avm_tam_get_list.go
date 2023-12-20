package tr064model

import "encoding/xml"

// GetListResponse AUTO-GENERATED (do not edit) model from [x_tamSCPD],
// based on SOAP action 'GetList', Fritz!Box-System-Version 164.07.57
//
// [x_tamSCPD]: http://fritz.box:49000/x_tamSCPD.xml
type GetListResponse struct {
	XMLName xml.Name // rather for debug purpose
	TAMList string   `xml:"NewTAMList"` //
}
