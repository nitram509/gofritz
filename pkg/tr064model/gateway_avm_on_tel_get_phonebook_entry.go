package tr064model

import "encoding/xml"

// GetPhonebookEntryResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'GetPhonebookEntry', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetPhonebookEntryResponse struct {
	XMLName            xml.Name // rather for debug purpose
	PhonebookEntryData string   `xml:"NewPhonebookEntryData"` //
}
