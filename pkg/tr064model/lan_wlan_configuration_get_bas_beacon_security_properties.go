package tr064model

import "encoding/xml"

// GetBasBeaconSecurityPropertiesResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetBasBeaconSecurityProperties', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetBasBeaconSecurityPropertiesResponse struct {
	XMLName                 xml.Name // rather for debug purpose
	BasicEncryptionModes    string   `xml:"NewBasicEncryptionModes"`    //
	BasicAuthenticationMode string   `xml:"NewBasicAuthenticationMode"` //
}
