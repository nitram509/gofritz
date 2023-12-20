package tr064model

import "encoding/xml"

// GetCallBarringEntryByNumResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'GetCallBarringEntryByNum', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetCallBarringEntryByNumResponse struct {
	XMLName            xml.Name // rather for debug purpose
	PhonebookEntryData string   `xml:"NewPhonebookEntryData"` //
}
