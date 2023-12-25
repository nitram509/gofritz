package tr064model

import "encoding/xml"

// GetSubnetMaskResponse AUTO-GENERATED (do not edit) model from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetSubnetMask', Fritz!Box-System-Version 141.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
type GetSubnetMaskResponse struct {
	XMLName    xml.Name // rather for debug purpose
	SubnetMask string   `xml:"NewSubnetMask"` //
}
