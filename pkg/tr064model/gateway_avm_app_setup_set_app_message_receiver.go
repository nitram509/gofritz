package tr064model

import "encoding/xml"

// SetAppMessageReceiverResponse AUTO-GENERATED (do not edit) model from [x_appsetupSCPD],
// based on SOAP action 'SetAppMessageReceiver', Fritz!Box-System-Version 164.08.00
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
type SetAppMessageReceiverResponse struct {
	XMLName          xml.Name // rather for debug purpose
	EncryptionSecret string   `xml:"EncryptionSecret"` //
	BoxSenderId      string   `xml:"BoxSenderId"`      //
}
