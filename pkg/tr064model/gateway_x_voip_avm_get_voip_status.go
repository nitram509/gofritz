package tr064model

import "encoding/xml"

// XavmGetVoIPStatusResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetVoIPStatus', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetVoIPStatusResponse struct {
	XMLName    xml.Name // rather for debug purpose
	VoIPStatus string   `xml:"NewX_AVM-DE_VoIPStatus"` // default=unknown; oneOf=[disabled,not registered,registered,connected,unknown]
}
