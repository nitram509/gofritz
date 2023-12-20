package tr064model

import "encoding/xml"

// XavmGetVoIPCommonAreaCodeResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetVoIPCommonAreaCode', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetVoIPCommonAreaCodeResponse struct {
	XMLName   xml.Name // rather for debug purpose
	OKZ       string   `xml:"NewX_AVM-DE_OKZ"`       //
	OKZPrefix string   `xml:"NewX_AVM-DE_OKZPrefix"` //
}
