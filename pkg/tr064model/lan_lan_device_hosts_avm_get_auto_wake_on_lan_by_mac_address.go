package tr064model

import "encoding/xml"

// XavmGetAutoWakeOnLANByMACAddressResponse AUTO-GENERATED (do not edit) model from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetAutoWakeOnLANByMACAddress', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
type XavmGetAutoWakeOnLANByMACAddressResponse struct {
	XMLName        xml.Name // rather for debug purpose
	AutoWOLEnabled bool     `xml:"NewAutoWOLEnabled"` // default=0
}
