package tr064model

// XavmGetSupportDataEnableResponse AUTO-GENERATED (do not edit) model from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetSupportDataEnable', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
type XavmGetSupportDataEnableResponse struct {
	SupportDataEnabled bool `xml:"NewX_AVM-DE_SupportDataEnabled"` // default=0
}
