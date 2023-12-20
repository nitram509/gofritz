package tr064model

// GetAddressRangeResponse AUTO-GENERATED (do not edit) model from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetAddressRange', Fritz!Box-System-Version 164.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
type GetAddressRangeResponse struct {
	MinAddress string `xml:"NewMinAddress"` //
	MaxAddress string `xml:"NewMaxAddress"` //
}
