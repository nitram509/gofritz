package tr064model

// GetHostNumberOfEntriesResponse AUTO-GENERATED (do not edit) model from [hostsSCPD],
// based on SOAP action 'GetHostNumberOfEntries', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
type GetHostNumberOfEntriesResponse struct {
	HostNumberOfEntries int `xml:"NewHostNumberOfEntries"` // default=0
}
