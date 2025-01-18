package tr064model

import "encoding/xml"

// GetStateResponse AUTO-GENERATED (do not edit) model from [x_authSCPD],
// based on SOAP action 'GetState', Fritz!Box-System-Version 164.08.00
//
// [x_authSCPD]: http://fritz.box:49000/x_authSCPD.xml
type GetStateResponse struct {
	XMLName xml.Name // rather for debug purpose
	State   string   `xml:"NewState"` //
}
