package tr064model

import "encoding/xml"

// GetHostNumberOfEntriesResponse AUTO-GENERATED (do not edit) model from [hostsSCPD],
// based on SOAP action 'GetHostNumberOfEntries', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
type GetHostNumberOfEntriesResponse struct {
	XMLName             xml.Name // rather for debug purpose
	HostNumberOfEntries int      `xml:"NewHostNumberOfEntries"` // default=0
}
