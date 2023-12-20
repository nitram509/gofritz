package tr064model

// GetChannelInfoResponse auto generated model from [wlanconfigSCPD],
// based on SOAP action 'GetChannelInfo', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetChannelInfoResponse struct {
	Channel            int    `xml:"NewChannel"`                     // default=0
	PossibleChannels   string `xml:"NewPossibleChannels"`            //
	AutoChannelEnabled bool   `xml:"NewX_AVM-DE_AutoChannelEnabled"` // default=0
	FrequencyBand      string `xml:"NewX_AVM-DE_FrequencyBand"`      // default=unknown; oneOf=[2400,5000,6000,unknown]
}
