package tr064model

import "encoding/xml"

// GetInfoByIndexResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'GetInfoByIndex', Fritz!Box-System-Version 164.08.00
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetInfoByIndexResponse struct {
	XMLName     xml.Name // rather for debug purpose
	Enable      bool     `xml:"NewEnable"`      // default=0
	Status      string   `xml:"NewStatus"`      //
	LastConnect string   `xml:"NewLastConnect"` //
	Url         string   `xml:"NewUrl"`         //
	ServiceId   string   `xml:"NewServiceId"`   //
	Username    string   `xml:"NewUsername"`    //
	Name        string   `xml:"NewName"`        //
}
