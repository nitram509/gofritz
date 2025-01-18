package tr064model

import "encoding/xml"

// GetExistingVoIPNumbersResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'GetExistingVoIPNumbers', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type GetExistingVoIPNumbersResponse struct {
	XMLName             xml.Name // rather for debug purpose
	ExistingVoIPNumbers int      `xml:"NewExistingVoIPNumbers"` // default=0
}
