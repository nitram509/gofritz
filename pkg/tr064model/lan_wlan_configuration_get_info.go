package tr064model

import "encoding/xml"

// GetWlanConfigurationInfoResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetWlanConfigurationInfoResponse struct {
	XMLName                  xml.Name // rather for debug purpose
	Enable                   bool     `xml:"NewEnable"`                       // default=0
	Status                   string   `xml:"NewStatus"`                       // oneOf=[Up,Disabled]
	MaxBitRate               string   `xml:"NewMaxBitRate"`                   //
	Channel                  int      `xml:"NewChannel"`                      // default=0
	SSID                     string   `xml:"NewSSID"`                         //
	BeaconType               string   `xml:"NewBeaconType"`                   // oneOf=[None,Basic,WPA,11i,WPAand11i,WPA3,11iandWPA3,OWE,OWETrans]
	PossibleBeaconTypes      string   `xml:"NewX_AVM-DE_PossibleBeaconTypes"` //
	MACAddressControlEnabled bool     `xml:"NewMACAddressControlEnabled"`     // default=0
	Standard                 string   `xml:"NewStandard"`                     // oneOf=[b,g,n,ac,ax,be,]
	BSSID                    string   `xml:"NewBSSID"`                        //
	BasicEncryptionModes     string   `xml:"NewBasicEncryptionModes"`         //
	BasicAuthenticationMode  string   `xml:"NewBasicAuthenticationMode"`      //
	MaxCharsSSID             int      `xml:"NewMaxCharsSSID"`                 // default=32
	MinCharsSSID             int      `xml:"NewMinCharsSSID"`                 // default=1
	AllowedCharsSSID         string   `xml:"NewAllowedCharsSSID"`             // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~
	MinCharsPSK              int      `xml:"NewMinCharsPSK"`                  // default=64
	MaxCharsPSK              int      `xml:"NewMaxCharsPSK"`                  // default=64
	AllowedCharsPSK          string   `xml:"NewAllowedCharsPSK"`              // default=0123456789ABCDEFabcdef
	FrequencyBand            string   `xml:"NewX_AVM-DE_FrequencyBand"`       // default=unknown; oneOf=[2400,5000,6000,unknown]
	WLANGlobalEnable         bool     `xml:"NewX_AVM-DE_WLANGlobalEnable"`    // default=0
}
