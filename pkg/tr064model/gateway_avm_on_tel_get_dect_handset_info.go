package tr064model

import "encoding/xml"

// GetDECTHandsetInfoResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'GetDECTHandsetInfo', Fritz!Box-System-Version 141.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetDECTHandsetInfoResponse struct {
	XMLName     xml.Name // rather for debug purpose
	HandsetName string   `xml:"NewHandsetName"` //
	PhonebookID int      `xml:"NewPhonebookID"` // default=0
}
