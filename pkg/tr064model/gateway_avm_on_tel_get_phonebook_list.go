package tr064model

import "encoding/xml"

// GetPhonebookListResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'GetPhonebookList', Fritz!Box-System-Version 141.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetPhonebookListResponse struct {
	XMLName       xml.Name // rather for debug purpose
	PhonebookList string   `xml:"NewPhonebookList"` //
}
