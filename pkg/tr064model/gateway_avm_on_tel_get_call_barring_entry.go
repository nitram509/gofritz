package tr064model

import "encoding/xml"

// GetCallBarringEntryResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'GetCallBarringEntry', Fritz!Box-System-Version 141.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetCallBarringEntryResponse struct {
	XMLName            xml.Name // rather for debug purpose
	PhonebookEntryData string   `xml:"NewPhonebookEntryData"` //
}
