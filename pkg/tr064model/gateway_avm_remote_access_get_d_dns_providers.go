package tr064model

import "encoding/xml"

// GetDDNSProvidersResponse AUTO-GENERATED (do not edit) model from [x_remoteSCPD],
// based on SOAP action 'GetDDNSProviders', Fritz!Box-System-Version 141.07.57
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
type GetDDNSProvidersResponse struct {
	XMLName      xml.Name // rather for debug purpose
	ProviderList string   `xml:"NewProviderList"` //
}
