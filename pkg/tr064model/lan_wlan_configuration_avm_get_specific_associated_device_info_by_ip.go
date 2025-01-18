package tr064model

import "encoding/xml"

// XavmGetSpecificAssociatedDeviceInfoByIpResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetSpecificAssociatedDeviceInfoByIp', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type XavmGetSpecificAssociatedDeviceInfoByIpResponse struct {
	XMLName                    xml.Name // rather for debug purpose
	AssociatedDeviceMACAddress string   `xml:"NewAssociatedDeviceMACAddress"` //
	AssociatedDeviceAuthState  bool     `xml:"NewAssociatedDeviceAuthState"`  // default=0
	Speed                      int      `xml:"NewX_AVM-DE_Speed"`             // default=0
	SignalStrength             int      `xml:"NewX_AVM-DE_SignalStrength"`    // default=0
	ChannelWidth               int      `xml:"NewX_AVM-DE_ChannelWidth"`      // default=0; oneOf=[0,20,40,80,160,320]
}
