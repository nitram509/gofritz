package tr064model

import "encoding/xml"

// XavmGetClientResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetClient', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetClientResponse struct {
	XMLName             xml.Name // rather for debug purpose
	ClientUsername      string   `xml:"NewX_AVM-DE_ClientUsername"`      //
	ClientRegistrar     string   `xml:"NewX_AVM-DE_ClientRegistrar"`     //
	ClientRegistrarPort int      `xml:"NewX_AVM-DE_ClientRegistrarPort"` // default=0
	PhoneName           string   `xml:"NewX_AVM-DE_PhoneName"`           //
	OutGoingNumber      string   `xml:"NewX_AVM-DE_OutGoingNumber"`      //
}
