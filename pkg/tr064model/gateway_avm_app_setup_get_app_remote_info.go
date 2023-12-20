package tr064model

// GetAppRemoteInfoResponse AUTO-GENERATED (do not edit) model from [x_appsetupSCPD],
// based on SOAP action 'GetAppRemoteInfo', Fritz!Box-System-Version 164.07.57
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
type GetAppRemoteInfoResponse struct {
	SubnetMask              string `xml:"NewSubnetMask"`              //
	IPAddress               string `xml:"NewIPAddress"`               //
	ExternalIPAddress       string `xml:"NewExternalIPAddress"`       //
	ExternalIPv6Address     string `xml:"NewExternalIPv6Address"`     //
	RemoteAccessDDNSEnabled bool   `xml:"NewRemoteAccessDDNSEnabled"` // default=0
	RemoteAccessDDNSDomain  string `xml:"NewRemoteAccessDDNSDomain"`  //
	MyFritzEnabled          bool   `xml:"NewMyFritzEnabled"`          // default=0
	MyFritzDynDNSName       string `xml:"NewMyFritzDynDNSName"`       //
}
