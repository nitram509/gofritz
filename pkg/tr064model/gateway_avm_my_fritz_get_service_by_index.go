package tr064model

// GetServiceByIndexResponse auto generated model from [x_myfritzSCPD],
// based on SOAP action 'GetServiceByIndex', Fritz!Box-System-Version 164.07.57
//
// [x_myfritzSCPD]: http://fritz.box:49000/x_myfritzSCPD.xml
type GetServiceByIndexResponse struct {
	Enabled               bool   `xml:"NewEnabled"`               // default=0
	Name                  string `xml:"NewName"`                  //
	Scheme                string `xml:"NewScheme"`                //
	Port                  int    `xml:"NewPort"`                  // default=0
	URLPath               string `xml:"NewURLPath"`               //
	Type                  string `xml:"NewType"`                  //
	IPv4ForwardingWarning int    `xml:"NewIPv4ForwardingWarning"` // default=0
	IPv4Addresses         string `xml:"NewIPv4Addresses"`         //
	IPv6Addresses         string `xml:"NewIPv6Addresses"`         //
	IPv6InterfaceIDs      string `xml:"NewIPv6InterfaceIDs"`      //
	MACAddress            string `xml:"NewMACAddress"`            //
	HostName              string `xml:"NewHostName"`              //
	DynDnsLabel           string `xml:"NewDynDnsLabel"`           //
	Status                int    `xml:"NewStatus"`                // default=0
}
