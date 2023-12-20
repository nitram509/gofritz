package tr064model

// GetPhonebookResponse auto generated model from [x_contactSCPD],
// based on SOAP action 'GetPhonebook', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetPhonebookResponse struct {
	PhonebookName    string `xml:"NewPhonebookName"`    //
	PhonebookExtraID string `xml:"NewPhonebookExtraID"` //
	PhonebookURL     string `xml:"NewPhonebookURL"`     //
}
