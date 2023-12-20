package tr064model

import "encoding/xml"

// NewFilelinkEntryResponse AUTO-GENERATED (do not edit) model from [x_filelinksSCPD],
// based on SOAP action 'NewFilelinkEntry', Fritz!Box-System-Version 164.07.57
//
// [x_filelinksSCPD]: http://fritz.box:49000/x_filelinksSCPD.xml
type NewFilelinkEntryResponse struct {
	XMLName xml.Name // rather for debug purpose
	ID      string   `xml:"NewID"` //
}
