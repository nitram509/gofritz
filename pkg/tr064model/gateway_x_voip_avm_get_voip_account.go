package tr064model

// XavmGetVoIPAccountResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetVoIPAccount', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetVoIPAccountResponse struct {
	VoIPRegistrar     string `xml:"NewVoIPRegistrar"`       //
	VoIPNumber        string `xml:"NewVoIPNumber"`          //
	VoIPUsername      string `xml:"NewVoIPUsername"`        //
	VoIPPassword      string `xml:"NewVoIPPassword"`        //
	VoIPOutboundProxy string `xml:"NewVoIPOutboundProxy"`   //
	VoIPSTUNServer    string `xml:"NewVoIPSTUNServer"`      //
	VoIPStatus        string `xml:"NewX_AVM-DE_VoIPStatus"` // default=unknown; oneOf=[disabled,not registered,registered,connected,unknown]
}
