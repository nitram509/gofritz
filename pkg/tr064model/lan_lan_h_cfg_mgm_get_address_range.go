package tr064model

import "encoding/xml"

// GetAddressRangeResponse AUTO-GENERATED (do not edit) model from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetAddressRange', Fritz!Box-System-Version 141.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
type GetAddressRangeResponse struct {
	XMLName    xml.Name // rather for debug purpose
	MinAddress string   `xml:"NewMinAddress"` //
	MaxAddress string   `xml:"NewMaxAddress"` //
}
