package tr064model

// XavmGetAnonymousLoginResponse auto generated model from [lanconfigsecuritySCPD],
// based on SOAP action 'X_AVM-DE_GetAnonymousLogin', Fritz!Box-System-Version 164.07.57
//
// [lanconfigsecuritySCPD]: http://fritz.box:49000/lanconfigsecuritySCPD.xml
type XavmGetAnonymousLoginResponse struct {
	AnonymousLoginEnabled bool `xml:"NewX_AVM-DE_AnonymousLoginEnabled"` // default=0
	ButtonLoginEnabled    bool `xml:"NewX_AVM-DE_ButtonLoginEnabled"`    // default=0
}
