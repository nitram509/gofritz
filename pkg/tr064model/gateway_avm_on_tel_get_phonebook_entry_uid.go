package tr064model

import "encoding/xml"

// GetPhonebookEntryUIDResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'GetPhonebookEntryUID', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetPhonebookEntryUIDResponse struct {
	XMLName            xml.Name // rather for debug purpose
	PhonebookEntryData string   `xml:"NewPhonebookEntryData"` //
}
