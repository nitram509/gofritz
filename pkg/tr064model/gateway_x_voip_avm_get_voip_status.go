package tr064model

// XavmGetVoIPStatusResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetVoIPStatus', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetVoIPStatusResponse struct {
	VoIPStatus string `xml:"NewX_AVM-DE_VoIPStatus"` // default=unknown; oneOf=[disabled,not registered,registered,connected,unknown]
}
