package tr064model

import "encoding/xml"

// XavmSetClient4Response AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_SetClient4', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmSetClient4Response struct {
	XMLName        xml.Name // rather for debug purpose
	InternalNumber string   `xml:"NewX_AVM-DE_InternalNumber"` //
}
