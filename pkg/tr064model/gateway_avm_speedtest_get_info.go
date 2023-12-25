package tr064model

import "encoding/xml"

// GetAvmSpeedtestInfoResponse AUTO-GENERATED (do not edit) model from [x_speedtestSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 141.07.57
//
// [x_speedtestSCPD]: http://fritz.box:49000/x_speedtestSCPD.xml
type GetAvmSpeedtestInfoResponse struct {
	XMLName           xml.Name // rather for debug purpose
	EnableTcp         bool     `xml:"NewEnableTcp"`         // default=0
	EnableUdp         bool     `xml:"NewEnableUdp"`         // default=0
	EnableUdpBidirect bool     `xml:"NewEnableUdpBidirect"` // default=0
	WANEnableTcp      bool     `xml:"NewWANEnableTcp"`      // default=0
	WANEnableUdp      bool     `xml:"NewWANEnableUdp"`      // default=0
	PortTcp           int      `xml:"NewPortTcp"`           // default=4711
	PortUdp           int      `xml:"NewPortUdp"`           // default=4711
	PortUdpBidirect   int      `xml:"NewPortUdpBidirect"`   // default=4712
}
