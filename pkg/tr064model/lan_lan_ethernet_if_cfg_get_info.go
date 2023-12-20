package tr064model

// GetLanEthernetIfCfgInfoResponse auto generated model from [ethifconfigSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [ethifconfigSCPD]: http://fritz.box:49000/ethifconfigSCPD.xml
type GetLanEthernetIfCfgInfoResponse struct {
	Enable     bool   `xml:"NewEnable"`     // default=0
	Status     string `xml:"NewStatus"`     //
	MACAddress string `xml:"NewMACAddress"` //
	MaxBitRate string `xml:"NewMaxBitRate"` //
	DuplexMode string `xml:"NewDuplexMode"` //
}
