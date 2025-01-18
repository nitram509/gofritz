package tr064model

import "encoding/xml"

// XavmGetClient3Response AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetClient3', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetClient3Response struct {
	XMLName                 xml.Name // rather for debug purpose
	ClientUsername          string   `xml:"NewX_AVM-DE_ClientUsername"`          //
	ClientRegistrar         string   `xml:"NewX_AVM-DE_ClientRegistrar"`         //
	ClientRegistrarPort     int      `xml:"NewX_AVM-DE_ClientRegistrarPort"`     // default=0
	PhoneName               string   `xml:"NewX_AVM-DE_PhoneName"`               //
	ClientId                string   `xml:"NewX_AVM-DE_ClientId"`                //
	OutGoingNumber          string   `xml:"NewX_AVM-DE_OutGoingNumber"`          //
	InComingNumbers         string   `xml:"NewX_AVM-DE_InComingNumbers"`         //
	ExternalRegistration    bool     `xml:"NewX_AVM-DE_ExternalRegistration"`    // default=0
	InternalNumber          string   `xml:"NewX_AVM-DE_InternalNumber"`          //
	DelayedCallNotification bool     `xml:"NewX_AVM-DE_DelayedCallNotification"` // default=0
}
