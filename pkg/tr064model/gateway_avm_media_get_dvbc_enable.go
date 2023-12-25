package tr064model

import "encoding/xml"

// GetDVBCEnableResponse AUTO-GENERATED (do not edit) model from [x_mediaSCPD],
// based on SOAP action 'GetDVBCEnable', Fritz!Box-System-Version 141.07.57
//
// [x_mediaSCPD]: http://fritz.box:49000/x_mediaSCPD.xml
type GetDVBCEnableResponse struct {
	XMLName     xml.Name // rather for debug purpose
	DVBCEnabled bool     `xml:"NewDVBCEnabled"` // default=0
}
