package tr064model

import "encoding/xml"

// SetPhonebookEntryUIDResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'SetPhonebookEntryUID', Fritz!Box-System-Version 141.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type SetPhonebookEntryUIDResponse struct {
	XMLName                xml.Name // rather for debug purpose
	PhonebookEntryUniqueID int      `xml:"NewPhonebookEntryUniqueID"` // default=0
}
