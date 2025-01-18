package tr064model

import "encoding/xml"

// GetTicketIDStatusResponse AUTO-GENERATED (do not edit) model from [x_hostfilterSCPD],
// based on SOAP action 'GetTicketIDStatus', Fritz!Box-System-Version 164.08.00
//
// [x_hostfilterSCPD]: http://fritz.box:49000/x_hostfilterSCPD.xml
type GetTicketIDStatusResponse struct {
	XMLName        xml.Name // rather for debug purpose
	TicketIDStatus string   `xml:"NewTicketIDStatus"` // oneOf=[marked,unmarked,invalid]
}
