package tr064model

// GetAvmTamInfoResponse AUTO-GENERATED (do not edit) model from [x_tamSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_tamSCPD]: http://fritz.box:49000/x_tamSCPD.xml
type GetAvmTamInfoResponse struct {
	Enable       bool   `xml:"NewEnable"`       // default=0
	Name         string `xml:"NewName"`         //
	TAMRunning   bool   `xml:"NewTAMRunning"`   // default=0
	Stick        int    `xml:"NewStick"`        // default=0
	Status       int    `xml:"NewStatus"`       // default=0
	Capacity     int    `xml:"NewCapacity"`     // default=0
	Mode         string `xml:"NewMode"`         // oneOf=[play_announcement,record_message,timeprofile]
	RingSeconds  int    `xml:"NewRingSeconds"`  // default=0
	PhoneNumbers string `xml:"NewPhoneNumbers"` //
}
