package tr064model

import "encoding/xml"

// GetLanEthernetIfCfgStatisticsResponse AUTO-GENERATED (do not edit) model from [ethifconfigSCPD],
// based on SOAP action 'GetStatistics', Fritz!Box-System-Version 164.08.00
//
// [ethifconfigSCPD]: http://fritz.box:49000/ethifconfigSCPD.xml
type GetLanEthernetIfCfgStatisticsResponse struct {
	XMLName         xml.Name // rather for debug purpose
	BytesSent       int      `xml:"NewBytesSent"`       // default=0
	BytesReceived   int      `xml:"NewBytesReceived"`   // default=0
	PacketsSent     int      `xml:"NewPacketsSent"`     // default=0
	PacketsReceived int      `xml:"NewPacketsReceived"` // default=0
}
