package tr064model

// SetAppMessageReceiverResponse AUTO-GENERATED (do not edit) model from [x_appsetupSCPD],
// based on SOAP action 'SetAppMessageReceiver', Fritz!Box-System-Version 164.07.57
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
type SetAppMessageReceiverResponse struct {
	EncryptionSecret string `xml:"EncryptionSecret"` //
	BoxSenderId      string `xml:"BoxSenderId"`      //
}
