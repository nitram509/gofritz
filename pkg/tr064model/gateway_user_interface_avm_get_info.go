package tr064model

import "encoding/xml"

// XavmGetUserInterfaceInfoResponse AUTO-GENERATED (do not edit) model from [userifSCPD],
// based on SOAP action 'X_AVM-DE_GetInfo', Fritz!Box-System-Version 141.07.57
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
type XavmGetUserInterfaceInfoResponse struct {
	XMLName          xml.Name // rather for debug purpose
	AutoUpdateMode   string   `xml:"NewX_AVM-DE_AutoUpdateMode"`   // oneOf=[off,all,important,check]
	UpdateTime       string   `xml:"NewX_AVM-DE_UpdateTime"`       // default=0001-01-01T00:00:00; type=Datetime
	LastFwVersion    string   `xml:"NewX_AVM-DE_LastFwVersion"`    //
	LastInfoUrl      string   `xml:"NewX_AVM-DE_LastInfoUrl"`      //
	CurrentFwVersion string   `xml:"NewX_AVM-DE_CurrentFwVersion"` //
	UpdateSuccessful string   `xml:"NewX_AVM-DE_UpdateSuccessful"` // oneOf=[unknown,failed,succeeded]
}
