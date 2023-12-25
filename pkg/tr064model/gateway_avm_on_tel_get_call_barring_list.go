package tr064model

import "encoding/xml"

// GetCallBarringListResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'GetCallBarringList', Fritz!Box-System-Version 141.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetCallBarringListResponse struct {
	XMLName      xml.Name // rather for debug purpose
	PhonebookURL string   `xml:"NewPhonebookURL"` //
}
