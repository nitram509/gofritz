package tr064model

import "encoding/xml"

// XavmGetNumberOfNumbersResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetNumberOfNumbers', Fritz!Box-System-Version 141.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetNumberOfNumbersResponse struct {
	XMLName         xml.Name // rather for debug purpose
	NumberOfNumbers int      `xml:"NewNumberOfNumbers"` // default=0
}
