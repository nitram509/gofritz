package tr064model

// GetTotalPacketsSentResponse AUTO-GENERATED (do not edit) model from [wancommonifconfigSCPD],
// based on SOAP action 'GetTotalPacketsSent', Fritz!Box-System-Version 164.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
type GetTotalPacketsSentResponse struct {
	TotalPacketsSent int `xml:"NewTotalPacketsSent"` // default=0
}
