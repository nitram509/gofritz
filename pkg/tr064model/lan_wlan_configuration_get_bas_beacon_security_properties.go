package tr064model

// GetBasBeaconSecurityPropertiesResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetBasBeaconSecurityProperties', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetBasBeaconSecurityPropertiesResponse struct {
	BasicEncryptionModes    string `xml:"NewBasicEncryptionModes"`    //
	BasicAuthenticationMode string `xml:"NewBasicAuthenticationMode"` //
}
