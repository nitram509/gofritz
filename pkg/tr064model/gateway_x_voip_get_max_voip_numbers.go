package tr064model

import "encoding/xml"

// GetMaxVoIPNumbersResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'GetMaxVoIPNumbers', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type GetMaxVoIPNumbersResponse struct {
	XMLName        xml.Name // rather for debug purpose
	MaxVoIPNumbers int      `xml:"NewMaxVoIPNumbers"` // default=0
}
