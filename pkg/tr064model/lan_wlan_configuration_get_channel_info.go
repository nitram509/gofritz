package tr064model

import "encoding/xml"

// GetChannelInfoResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetChannelInfo', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetChannelInfoResponse struct {
	XMLName            xml.Name // rather for debug purpose
	Channel            int      `xml:"NewChannel"`                     // default=0
	PossibleChannels   string   `xml:"NewPossibleChannels"`            //
	AutoChannelEnabled bool     `xml:"NewX_AVM-DE_AutoChannelEnabled"` // default=0
	FrequencyBand      string   `xml:"NewX_AVM-DE_FrequencyBand"`      // default=unknown; oneOf=[2400,5000,6000,unknown]
}
