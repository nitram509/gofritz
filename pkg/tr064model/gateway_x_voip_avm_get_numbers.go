package tr064model

import "encoding/xml"

// XavmGetNumbersResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetNumbers', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetNumbersResponse struct {
	XMLName    xml.Name // rather for debug purpose
	NumberList string   `xml:"NewNumberList"` //
}
