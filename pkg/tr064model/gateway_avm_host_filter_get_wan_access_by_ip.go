package tr064model

import "encoding/xml"

// GetWANAccessByIPResponse AUTO-GENERATED (do not edit) model from [x_hostfilterSCPD],
// based on SOAP action 'GetWANAccessByIP', Fritz!Box-System-Version 164.07.57
//
// [x_hostfilterSCPD]: http://fritz.box:49000/x_hostfilterSCPD.xml
type GetWANAccessByIPResponse struct {
	XMLName   xml.Name // rather for debug purpose
	Disallow  bool     `xml:"NewDisallow"`  // default=0
	WANAccess string   `xml:"NewWANAccess"` // oneOf=[error,denied,granted]
}
