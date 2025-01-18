package tr064model

import "encoding/xml"

// GetLanEthernetIfCfgInfoResponse AUTO-GENERATED (do not edit) model from [ethifconfigSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [ethifconfigSCPD]: http://fritz.box:49000/ethifconfigSCPD.xml
type GetLanEthernetIfCfgInfoResponse struct {
	XMLName    xml.Name // rather for debug purpose
	Enable     bool     `xml:"NewEnable"`     // default=0
	Status     string   `xml:"NewStatus"`     //
	MACAddress string   `xml:"NewMACAddress"` //
	MaxBitRate string   `xml:"NewMaxBitRate"` //
	DuplexMode string   `xml:"NewDuplexMode"` //
}
