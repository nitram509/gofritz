package tr064model

// XavmGetClientByClientIdResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetClientByClientId', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetClientByClientIdResponse struct {
	ClientId                string `xml:"NewX_AVM-DE_ClientId"`                //
	ClientIndex             int    `xml:"NewX_AVM-DE_ClientIndex"`             // default=0
	ClientUsername          string `xml:"NewX_AVM-DE_ClientUsername"`          //
	ClientRegistrar         string `xml:"NewX_AVM-DE_ClientRegistrar"`         //
	ClientRegistrarPort     int    `xml:"NewX_AVM-DE_ClientRegistrarPort"`     // default=0
	PhoneName               string `xml:"NewX_AVM-DE_PhoneName"`               //
	OutGoingNumber          string `xml:"NewX_AVM-DE_OutGoingNumber"`          //
	InComingNumbers         string `xml:"NewX_AVM-DE_InComingNumbers"`         //
	ExternalRegistration    bool   `xml:"NewX_AVM-DE_ExternalRegistration"`    // default=0
	InternalNumber          string `xml:"NewX_AVM-DE_InternalNumber"`          //
	DelayedCallNotification bool   `xml:"NewX_AVM-DE_DelayedCallNotification"` // default=0
}
