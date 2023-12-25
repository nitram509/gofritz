package tr064model

import "encoding/xml"

// GetUSPControllerNumberOfEntriesResponse AUTO-GENERATED (do not edit) model from [x_uspcontrollerSCPD],
// based on SOAP action 'GetUSPControllerNumberOfEntries', Fritz!Box-System-Version 141.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
type GetUSPControllerNumberOfEntriesResponse struct {
	XMLName                      xml.Name // rather for debug purpose
	USPControllerNumberOfEntries int      `xml:"NewUSPControllerNumberOfEntries"` // default=0
}
