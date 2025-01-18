package tr064model

import "encoding/xml"

// XavmGetWLANHybridModeResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetWLANHybridMode', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type XavmGetWLANHybridModeResponse struct {
	XMLName       xml.Name // rather for debug purpose
	Enable        bool     `xml:"NewEnable"`        // default=0
	BeaconType    string   `xml:"NewBeaconType"`    // oneOf=[None,Basic,WPA,11i,WPAand11i,WPA3,11iandWPA3,OWE,OWETrans]
	KeyPassphrase string   `xml:"NewKeyPassphrase"` //
	SSID          string   `xml:"NewSSID"`          //
	BSSID         string   `xml:"NewBSSID"`         //
	TrafficMode   string   `xml:"NewTrafficMode"`   //
	ManualSpeed   bool     `xml:"NewManualSpeed"`   // default=0
	MaxSpeedDS    int      `xml:"NewMaxSpeedDS"`    // default=0
	MaxSpeedUS    int      `xml:"NewMaxSpeedUS"`    // default=0
}
