package tr064model

import "encoding/xml"

// GetBoxSenderIdResponse AUTO-GENERATED (do not edit) model from [x_appsetupSCPD],
// based on SOAP action 'GetBoxSenderId', Fritz!Box-System-Version 164.07.57
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
type GetBoxSenderIdResponse struct {
	XMLName     xml.Name // rather for debug purpose
	BoxSenderId string   `xml:"NewBoxSenderId"` //
}
