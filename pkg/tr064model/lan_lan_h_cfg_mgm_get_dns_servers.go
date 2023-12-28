package tr064model

import "encoding/xml"

// GetDNSServersResponse AUTO-GENERATED (do not edit) model from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetDNSServers', Fritz!Box-System-Version 164.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
type GetDNSServersResponse struct {
	XMLName    xml.Name // rather for debug purpose
	DNSServers string   `xml:"NewDNSServers"` //
}
