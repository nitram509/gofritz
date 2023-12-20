package tr064model

import "encoding/xml"

// GetUSPMyFRITZEnableResponse AUTO-GENERATED (do not edit) model from [x_uspcontrollerSCPD],
// based on SOAP action 'GetUSPMyFRITZEnable', Fritz!Box-System-Version 164.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
type GetUSPMyFRITZEnableResponse struct {
	XMLName           xml.Name // rather for debug purpose
	USPMyFRITZEnabled bool     `xml:"NewUSPMyFRITZEnabled"` // default=0
}
