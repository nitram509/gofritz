package tr064model

import "encoding/xml"

// SetCallBarringEntryResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'SetCallBarringEntry', Fritz!Box-System-Version 164.08.00
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type SetCallBarringEntryResponse struct {
	XMLName                xml.Name // rather for debug purpose
	PhonebookEntryUniqueID int      `xml:"NewPhonebookEntryUniqueID"` // default=0
}
