package tr064model

import "encoding/xml"

// GetSpecificAssociatedDeviceInfoResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetSpecificAssociatedDeviceInfo', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetSpecificAssociatedDeviceInfoResponse struct {
	XMLName                   xml.Name // rather for debug purpose
	AssociatedDeviceIPAddress string   `xml:"NewAssociatedDeviceIPAddress"` //
	AssociatedDeviceAuthState bool     `xml:"NewAssociatedDeviceAuthState"` // default=0
	Speed                     int      `xml:"NewX_AVM-DE_Speed"`            // default=0
	SignalStrength            int      `xml:"NewX_AVM-DE_SignalStrength"`   // default=0
	ChannelWidth              int      `xml:"NewX_AVM-DE_ChannelWidth"`     // default=0
}
