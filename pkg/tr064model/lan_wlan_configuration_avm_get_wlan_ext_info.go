package tr064model

import "encoding/xml"

// XavmGetWLANExtInfoResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetWLANExtInfo', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type XavmGetWLANExtInfoResponse struct {
	XMLName          xml.Name // rather for debug purpose
	APEnabled        string   `xml:"NewX_AVM-DE_APEnabled"`        //
	APType           string   `xml:"NewX_AVM-DE_APType"`           //
	FrequencyBand    string   `xml:"NewX_AVM-DE_FrequencyBand"`    // default=unknown; oneOf=[2400,5000,6000,unknown]
	TimeoutActive    string   `xml:"NewX_AVM-DE_TimeoutActive"`    //
	Timeout          string   `xml:"NewX_AVM-DE_Timeout"`          //
	TimeRemain       string   `xml:"NewX_AVM-DE_TimeRemain"`       //
	NoForcedOff      string   `xml:"NewX_AVM-DE_NoForcedOff"`      //
	UserIsolation    string   `xml:"NewX_AVM-DE_UserIsolation"`    //
	EncryptionMode   string   `xml:"NewX_AVM-DE_EncryptionMode"`   //
	LastChangedStamp int      `xml:"NewX_AVM-DE_LastChangedStamp"` // default=0
}
