package tr064model

import "encoding/xml"

// GetNumberOfEntriesResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'GetNumberOfEntries', Fritz!Box-System-Version 141.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetNumberOfEntriesResponse struct {
	XMLName              xml.Name // rather for debug purpose
	OnTelNumberOfEntries int      `xml:"NewOnTelNumberOfEntries"` // default=0
}
