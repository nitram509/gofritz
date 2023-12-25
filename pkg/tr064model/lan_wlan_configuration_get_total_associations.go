package tr064model

import "encoding/xml"

// GetTotalAssociationsResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetTotalAssociations', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetTotalAssociationsResponse struct {
	XMLName           xml.Name // rather for debug purpose
	TotalAssociations int      `xml:"NewTotalAssociations"` // default=0
}
