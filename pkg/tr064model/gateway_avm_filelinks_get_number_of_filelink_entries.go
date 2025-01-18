package tr064model

import "encoding/xml"

// GetNumberOfFilelinkEntriesResponse AUTO-GENERATED (do not edit) model from [x_filelinksSCPD],
// based on SOAP action 'GetNumberOfFilelinkEntries', Fritz!Box-System-Version 164.08.00
//
// [x_filelinksSCPD]: http://fritz.box:49000/x_filelinksSCPD.xml
type GetNumberOfFilelinkEntriesResponse struct {
	XMLName         xml.Name // rather for debug purpose
	NumberOfEntries int      `xml:"NewNumberOfEntries"` // default=0
}
