package tr064model

// GetLanHCfgMgmInfoResponse auto generated model from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
type GetLanHCfgMgmInfoResponse struct {
	DHCPServerConfigurable bool   `xml:"NewDHCPServerConfigurable"` // default=0
	DHCPRelay              bool   `xml:"NewDHCPRelay"`              // default=0
	MinAddress             string `xml:"NewMinAddress"`             //
	MaxAddress             string `xml:"NewMaxAddress"`             //
	ReservedAddresses      string `xml:"NewReservedAddresses"`      //
	DHCPServerEnable       bool   `xml:"NewDHCPServerEnable"`       // default=0
	DNSServers             string `xml:"NewDNSServers"`             //
	DomainName             string `xml:"NewDomainName"`             //
	IPRouters              string `xml:"NewIPRouters"`              //
	SubnetMask             string `xml:"NewSubnetMask"`             //
}
