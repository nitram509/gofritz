package tr064model

// XavmGetWLANConnectionInfoResponse auto generated model from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetWLANConnectionInfo', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type XavmGetWLANConnectionInfoResponse struct {
	AssociatedDeviceMACAddress string `xml:"NewAssociatedDeviceMACAddress"`  //
	SSID                       string `xml:"NewSSID"`                        //
	BSSID                      string `xml:"NewBSSID"`                       //
	BeaconType                 string `xml:"NewBeaconType"`                  // oneOf=[none,Basic,WPA,11i,WPAand11i,WPA3,11iandWPA3,OWE,OWETrans]
	Channel                    int    `xml:"NewChannel"`                     // default=0
	Standard                   string `xml:"NewStandard"`                    // oneOf=[b,g,n,ac,ax,]
	AutoChannelEnabled         bool   `xml:"NewX_AVM-DE_AutoChannelEnabled"` // default=0
	ChannelWidth               int    `xml:"NewX_AVM-DE_ChannelWidth"`       // default=0
	FrequencyBand              string `xml:"NewX_AVM-DE_FrequencyBand"`      // default=unknown; oneOf=[2400,5000,6000,unknown]
	SignalStrength             int    `xml:"NewX_AVM-DE_SignalStrength"`     // default=0
	Speed                      int    `xml:"NewX_AVM-DE_Speed"`              // default=0
	SpeedRX                    int    `xml:"NewX_AVM-DE_SpeedRX"`            // default=0
	SpeedMax                   int    `xml:"NewX_AVM-DE_SpeedMax"`           // default=0
	SpeedRXMax                 int    `xml:"NewX_AVM-DE_SpeedRXMax"`         // default=0
}
