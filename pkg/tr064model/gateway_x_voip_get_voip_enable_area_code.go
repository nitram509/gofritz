package tr064model

// GetVoIPEnableAreaCodeResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'GetVoIPEnableAreaCode', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type GetVoIPEnableAreaCodeResponse struct {
	VoIPEnableAreaCode bool `xml:"NewVoIPEnableAreaCode"` // default=0
}
