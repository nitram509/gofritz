package tr064model

import "encoding/xml"

// XavmGetAnonymousLoginResponse AUTO-GENERATED (do not edit) model from [lanconfigsecuritySCPD],
// based on SOAP action 'X_AVM-DE_GetAnonymousLogin', Fritz!Box-System-Version 141.07.57
//
// [lanconfigsecuritySCPD]: http://fritz.box:49000/lanconfigsecuritySCPD.xml
type XavmGetAnonymousLoginResponse struct {
	XMLName               xml.Name // rather for debug purpose
	AnonymousLoginEnabled bool     `xml:"NewX_AVM-DE_AnonymousLoginEnabled"` // default=0
	ButtonLoginEnabled    bool     `xml:"NewX_AVM-DE_ButtonLoginEnabled"`    // default=0
}
