package tr064model

// GetSpecificHostEntryResponse auto generated model from [hostsSCPD],
// based on SOAP action 'GetSpecificHostEntry', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
type GetSpecificHostEntryResponse struct {
	IPAddress          string `xml:"NewIPAddress"`          //
	AddressSource      string `xml:"NewAddressSource"`      //
	LeaseTimeRemaining int    `xml:"NewLeaseTimeRemaining"` // default=0
	InterfaceType      string `xml:"NewInterfaceType"`      //
	Active             bool   `xml:"NewActive"`             // default=0
	HostName           string `xml:"NewHostName"`           //
}
