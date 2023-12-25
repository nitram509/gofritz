package tr064model

import "encoding/xml"

// GetNumberOfServicesResponse AUTO-GENERATED (do not edit) model from [x_myfritzSCPD],
// based on SOAP action 'GetNumberOfServices', Fritz!Box-System-Version 141.07.57
//
// [x_myfritzSCPD]: http://fritz.box:49000/x_myfritzSCPD.xml
type GetNumberOfServicesResponse struct {
	XMLName          xml.Name // rather for debug purpose
	NumberOfServices int      `xml:"NewNumberOfServices"` // default=0
}
