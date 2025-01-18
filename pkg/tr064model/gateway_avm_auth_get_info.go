package tr064model

import "encoding/xml"

// GetAvmAuthInfoResponse AUTO-GENERATED (do not edit) model from [x_authSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [x_authSCPD]: http://fritz.box:49000/x_authSCPD.xml
type GetAvmAuthInfoResponse struct {
	XMLName xml.Name // rather for debug purpose
	Enabled bool     `xml:"NewEnabled"` // default=0
}
