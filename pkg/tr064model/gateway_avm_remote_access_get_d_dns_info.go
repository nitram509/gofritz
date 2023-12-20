package tr064model

// GetDDNSInfoResponse auto generated model from [x_remoteSCPD],
// based on SOAP action 'GetDDNSInfo', Fritz!Box-System-Version 164.07.57
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
type GetDDNSInfoResponse struct {
	Enabled      bool   `xml:"NewEnabled"`      // default=0
	ProviderName string `xml:"NewProviderName"` //
	UpdateURL    string `xml:"NewUpdateURL"`    //
	Domain       string `xml:"NewDomain"`       //
	StatusIPv4   string `xml:"NewStatusIPv4"`   //
	StatusIPv6   string `xml:"NewStatusIPv6"`   //
	Username     string `xml:"NewUsername"`     //
	Mode         string `xml:"NewMode"`         // default=ddns_v4; oneOf=[ddns_v4,ddns_v6,ddns_both,ddns_together]
	ServerIPv4   string `xml:"NewServerIPv4"`   //
	ServerIPv6   string `xml:"NewServerIPv6"`   //
}
