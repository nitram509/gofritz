package tr064model

import "encoding/xml"

// XavmGetSupportDataEnableResponse AUTO-GENERATED (do not edit) model from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetSupportDataEnable', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
type XavmGetSupportDataEnableResponse struct {
	XMLName            xml.Name // rather for debug purpose
	SupportDataEnabled bool     `xml:"NewX_AVM-DE_SupportDataEnabled"` // default=0
}
