package tr064model

// GetTicketIDStatusResponse AUTO-GENERATED (do not edit) model from [x_hostfilterSCPD],
// based on SOAP action 'GetTicketIDStatus', Fritz!Box-System-Version 164.07.57
//
// [x_hostfilterSCPD]: http://fritz.box:49000/x_hostfilterSCPD.xml
type GetTicketIDStatusResponse struct {
	TicketIDStatus string `xml:"NewTicketIDStatus"` // oneOf=[marked,unmarked,invalid]
}
