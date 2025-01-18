package tr064model

import "encoding/xml"

// GetSSIDResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetSSID', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetSSIDResponse struct {
	XMLName xml.Name // rather for debug purpose
	SSID    string   `xml:"NewSSID"` //
}
