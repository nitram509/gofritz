package tr064model

import "encoding/xml"

// GetNumberOfDectEntriesResponse AUTO-GENERATED (do not edit) model from [x_dectSCPD],
// based on SOAP action 'GetNumberOfDectEntries', Fritz!Box-System-Version 141.07.57
//
// [x_dectSCPD]: http://fritz.box:49000/x_dectSCPD.xml
type GetNumberOfDectEntriesResponse struct {
	XMLName         xml.Name // rather for debug purpose
	NumberOfEntries int      `xml:"NewNumberOfEntries"` // default=0
}
