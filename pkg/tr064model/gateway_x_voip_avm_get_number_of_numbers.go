package tr064model

// XavmGetNumberOfNumbersResponse auto generated model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetNumberOfNumbers', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetNumberOfNumbersResponse struct {
	NumberOfNumbers int `xml:"NewNumberOfNumbers"` // default=0
}
