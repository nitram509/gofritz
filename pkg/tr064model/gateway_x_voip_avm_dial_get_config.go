package tr064model

import "encoding/xml"

// XavmDialGetConfigResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_DialGetConfig', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmDialGetConfigResponse struct {
	XMLName   xml.Name // rather for debug purpose
	PhoneName string   `xml:"NewX_AVM-DE_PhoneName"` //
}
