package tr064model

import "encoding/xml"

// XavmGetVoIPCommonCountryCodeResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetVoIPCommonCountryCode', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetVoIPCommonCountryCodeResponse struct {
	XMLName   xml.Name // rather for debug purpose
	LKZ       string   `xml:"NewX_AVM-DE_LKZ"`       //
	LKZPrefix string   `xml:"NewX_AVM-DE_LKZPrefix"` //
}
