package tr064model

import "encoding/xml"

// GetAvmUspControllerInfoResponse AUTO-GENERATED (do not edit) model from [x_uspcontrollerSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 141.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
type GetAvmUspControllerInfoResponse struct {
	XMLName                xml.Name // rather for debug purpose
	MinCharsEndpointID     int      `xml:"NewMinCharsEndpointID"`     // default=0
	MaxCharsEndpointID     int      `xml:"NewMaxCharsEndpointID"`     // default=64
	AllowedCharsEndpointID string   `xml:"NewAllowedCharsEndpointID"` // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz:-._%
	MinCharsHostname       int      `xml:"NewMinCharsHostname"`       // default=0
	MaxCharsHostname       int      `xml:"NewMaxCharsHostname"`       // default=255
	MinCharsPath           int      `xml:"NewMinCharsPath"`           // default=0
	MaxCharsPath           int      `xml:"NewMaxCharsPath"`           // default=1024
	MinCharsUsername       int      `xml:"NewMinCharsUsername"`       // default=0
	MaxCharsUsername       int      `xml:"NewMaxCharsUsername"`       // default=255
	MinCharsPassword       int      `xml:"NewMinCharsPassword"`       // default=0
	MaxCharsPassword       int      `xml:"NewMaxCharsPassword"`       // default=255
}
