package tr064model

import "encoding/xml"

// GetBSSIDResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetBSSID', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetBSSIDResponse struct {
	XMLName xml.Name // rather for debug purpose
	BSSID   string   `xml:"NewBSSID"` //
}
