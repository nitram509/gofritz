package tr064model

// GetSpecificAssociatedDeviceInfoResponse auto generated model from [wlanconfigSCPD],
// based on SOAP action 'GetSpecificAssociatedDeviceInfo', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetSpecificAssociatedDeviceInfoResponse struct {
	AssociatedDeviceIPAddress string `xml:"NewAssociatedDeviceIPAddress"` //
	AssociatedDeviceAuthState bool   `xml:"NewAssociatedDeviceAuthState"` // default=0
	Speed                     int    `xml:"NewX_AVM-DE_Speed"`            // default=0
	SignalStrength            int    `xml:"NewX_AVM-DE_SignalStrength"`   // default=0
	ChannelWidth              int    `xml:"NewX_AVM-DE_ChannelWidth"`     // default=0
}