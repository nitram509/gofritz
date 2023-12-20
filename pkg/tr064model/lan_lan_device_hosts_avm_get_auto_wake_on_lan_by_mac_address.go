package tr064model

// XavmGetAutoWakeOnLANByMACAddressResponse auto generated model from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetAutoWakeOnLANByMACAddress', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
type XavmGetAutoWakeOnLANByMACAddressResponse struct {
	AutoWOLEnabled bool `xml:"NewAutoWOLEnabled"` // default=0
}
