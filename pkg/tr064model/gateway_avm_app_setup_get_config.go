package tr064model

import "encoding/xml"

// GetConfigResponse AUTO-GENERATED (do not edit) model from [x_appsetupSCPD],
// based on SOAP action 'GetConfig', Fritz!Box-System-Version 164.07.57
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
type GetConfigResponse struct {
	XMLName            xml.Name // rather for debug purpose
	ConfigRight        string   `xml:"NewConfigRight"`        // default=NO; oneOf=[NO,RO,RW,UNDEFINED]
	AppRight           string   `xml:"NewAppRight"`           // default=NO; oneOf=[NO,RO,RW,UNDEFINED]
	NasRight           string   `xml:"NewNasRight"`           // default=NO; oneOf=[NO,RO,RW,UNDEFINED]
	PhoneRight         string   `xml:"NewPhoneRight"`         // default=NO; oneOf=[NO,RO,RW,UNDEFINED]
	DialRight          string   `xml:"NewDialRight"`          // default=NO; oneOf=[NO,RO,RW,UNDEFINED]
	HomeautoRight      string   `xml:"NewHomeautoRight"`      // default=NO; oneOf=[NO,RO,RW,UNDEFINED]
	InternetRights     bool     `xml:"NewInternetRights"`     // default=0
	AccessFromInternet bool     `xml:"NewAccessFromInternet"` // default=0
}
