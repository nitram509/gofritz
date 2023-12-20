package tr064model

import "encoding/xml"

// GetPhonebookResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'GetPhonebook', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetPhonebookResponse struct {
	XMLName          xml.Name // rather for debug purpose
	PhonebookName    string   `xml:"NewPhonebookName"`    //
	PhonebookExtraID string   `xml:"NewPhonebookExtraID"` //
	PhonebookURL     string   `xml:"NewPhonebookURL"`     //
}
