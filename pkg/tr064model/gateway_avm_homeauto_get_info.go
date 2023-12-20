package tr064model

// GetAvmHomeautoInfoResponse AUTO-GENERATED (do not edit) model from [x_homeautoSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_homeautoSCPD]: http://fritz.box:49000/x_homeautoSCPD.xml
type GetAvmHomeautoInfoResponse struct {
	AllowedCharsAIN    string `xml:"NewAllowedCharsAIN"`    // default=0123456789ABCDEFabcdef :-grptmp
	MaxCharsAIN        int    `xml:"NewMaxCharsAIN"`        // default=19
	MinCharsAIN        int    `xml:"NewMinCharsAIN"`        // default=1
	MaxCharsDeviceName int    `xml:"NewMaxCharsDeviceName"` // default=79
	MinCharsDeviceName int    `xml:"NewMinCharsDeviceName"` // default=1
}
