package tr064model

import "encoding/xml"

// XavmGetWLANConnectionInfoResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetWLANConnectionInfo', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type XavmGetWLANConnectionInfoResponse struct {
	XMLName                    xml.Name // rather for debug purpose
	AssociatedDeviceMACAddress string   `xml:"NewAssociatedDeviceMACAddress"`  //
	SSID                       string   `xml:"NewSSID"`                        //
	BSSID                      string   `xml:"NewBSSID"`                       //
	BeaconType                 string   `xml:"NewBeaconType"`                  // oneOf=[None,Basic,WPA,11i,WPAand11i,WPA3,11iandWPA3,OWE,OWETrans]
	Channel                    int      `xml:"NewChannel"`                     // default=0
	Standard                   string   `xml:"NewStandard"`                    // oneOf=[b,g,n,ac,ax,be,]
	AutoChannelEnabled         bool     `xml:"NewX_AVM-DE_AutoChannelEnabled"` // default=0
	ChannelWidth               int      `xml:"NewX_AVM-DE_ChannelWidth"`       // default=0; oneOf=[0,20,40,80,160,320]
	FrequencyBand              string   `xml:"NewX_AVM-DE_FrequencyBand"`      // default=unknown; oneOf=[2400,5000,6000,unknown]
	SignalStrength             int      `xml:"NewX_AVM-DE_SignalStrength"`     // default=0
	Speed                      int      `xml:"NewX_AVM-DE_Speed"`              // default=0
	SpeedRX                    int      `xml:"NewX_AVM-DE_SpeedRX"`            // default=0
	SpeedMax                   int      `xml:"NewX_AVM-DE_SpeedMax"`           // default=0
	SpeedRXMax                 int      `xml:"NewX_AVM-DE_SpeedRXMax"`         // default=0
}
