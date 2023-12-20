package tr064model

import "encoding/xml"

// GetCommonLinkPropertiesResponse AUTO-GENERATED (do not edit) model from [wancommonifconfigSCPD],
// based on SOAP action 'GetCommonLinkProperties', Fritz!Box-System-Version 164.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
type GetCommonLinkPropertiesResponse struct {
	XMLName                      xml.Name // rather for debug purpose
	WANAccessType                string   `xml:"NewWANAccessType"`                         // default=unknown; oneOf=[DSL,Ethernet,X_AVM-DE_Fiber,X_AVM-DE_UMTS,X_AVM-DE_Cable,X_AVM-DE_LTE,unknown]
	Layer1UpstreamMaxBitRate     int      `xml:"NewLayer1UpstreamMaxBitRate"`              // default=0
	Layer1DownstreamMaxBitRate   int      `xml:"NewLayer1DownstreamMaxBitRate"`            // default=0
	PhysicalLinkStatus           string   `xml:"NewPhysicalLinkStatus"`                    //
	DownstreamCurrentUtilization string   `xml:"NewX_AVM-DE_DownstreamCurrentUtilization"` //
	UpstreamCurrentUtilization   string   `xml:"NewX_AVM-DE_UpstreamCurrentUtilization"`   //
	DownstreamCurrentMaxSpeed    int      `xml:"NewX_AVM-DE_DownstreamCurrentMaxSpeed"`    // default=0
	UpstreamCurrentMaxSpeed      int      `xml:"NewX_AVM-DE_UpstreamCurrentMaxSpeed"`      // default=0
}
