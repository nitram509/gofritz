package tr064model

import "encoding/xml"

// MarkTicketResponse AUTO-GENERATED (do not edit) model from [x_hostfilterSCPD],
// based on SOAP action 'MarkTicket', Fritz!Box-System-Version 164.08.00
//
// [x_hostfilterSCPD]: http://fritz.box:49000/x_hostfilterSCPD.xml
type MarkTicketResponse struct {
	XMLName  xml.Name // rather for debug purpose
	TicketID string   `xml:"NewTicketID"` //
}
