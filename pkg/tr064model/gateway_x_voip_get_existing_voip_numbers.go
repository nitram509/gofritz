package tr064model

// GetExistingVoIPNumbersResponse auto generated model from [x_voipSCPD],
// based on SOAP action 'GetExistingVoIPNumbers', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type GetExistingVoIPNumbersResponse struct {
	ExistingVoIPNumbers int `xml:"NewExistingVoIPNumbers"` // default=0
}
