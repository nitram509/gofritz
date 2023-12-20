package tr064model

// GetMaxVoIPNumbersResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'GetMaxVoIPNumbers', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type GetMaxVoIPNumbersResponse struct {
	MaxVoIPNumbers int `xml:"NewMaxVoIPNumbers"` // default=0
}
