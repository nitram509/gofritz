package tr064model

// GetDECTHandsetInfoResponse auto generated model from [x_contactSCPD],
// based on SOAP action 'GetDECTHandsetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetDECTHandsetInfoResponse struct {
	HandsetName string `xml:"NewHandsetName"` //
	PhonebookID int    `xml:"NewPhonebookID"` // default=0
}
