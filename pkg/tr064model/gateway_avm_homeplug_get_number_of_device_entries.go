package tr064model

import "encoding/xml"

// GetNumberOfDeviceEntriesResponse AUTO-GENERATED (do not edit) model from [x_homeplugSCPD],
// based on SOAP action 'GetNumberOfDeviceEntries', Fritz!Box-System-Version 141.07.57
//
// [x_homeplugSCPD]: http://fritz.box:49000/x_homeplugSCPD.xml
type GetNumberOfDeviceEntriesResponse struct {
	XMLName         xml.Name // rather for debug purpose
	NumberOfEntries int      `xml:"NewNumberOfEntries"` // default=0
}
