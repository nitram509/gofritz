package tr064model

import "encoding/xml"

// GetIPRoutersListResponse AUTO-GENERATED (do not edit) model from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetIPRoutersList', Fritz!Box-System-Version 164.08.00
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
type GetIPRoutersListResponse struct {
	XMLName   xml.Name // rather for debug purpose
	IPRouters string   `xml:"NewIPRouters"` //
}
