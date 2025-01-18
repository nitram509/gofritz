package tr064model

import "encoding/xml"

// GetVoIPCommonAreaCodeResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'GetVoIPCommonAreaCode', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type GetVoIPCommonAreaCodeResponse struct {
	XMLName      xml.Name // rather for debug purpose
	VoIPAreaCode string   `xml:"NewVoIPAreaCode"` //
}
