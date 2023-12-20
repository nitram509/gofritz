package tr064model

// GetUSPMyFRITZEnableResponse auto generated model from [x_uspcontrollerSCPD],
// based on SOAP action 'GetUSPMyFRITZEnable', Fritz!Box-System-Version 164.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
type GetUSPMyFRITZEnableResponse struct {
	USPMyFRITZEnabled bool `xml:"NewUSPMyFRITZEnabled"` // default=0
}
