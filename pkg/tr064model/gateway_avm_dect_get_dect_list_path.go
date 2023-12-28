package tr064model

import "encoding/xml"

// GetDectListPathResponse AUTO-GENERATED (do not edit) model from [x_dectSCPD],
// based on SOAP action 'GetDectListPath', Fritz!Box-System-Version 164.07.57
//
// [x_dectSCPD]: http://fritz.box:49000/x_dectSCPD.xml
type GetDectListPathResponse struct {
	XMLName      xml.Name // rather for debug purpose
	DectListPath string   `xml:"NewDectListPath"` //
}
