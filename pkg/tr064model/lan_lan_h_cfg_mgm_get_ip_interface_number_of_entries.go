package tr064model

import "encoding/xml"

// GetIPInterfaceNumberOfEntriesResponse AUTO-GENERATED (do not edit) model from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetIPInterfaceNumberOfEntries', Fritz!Box-System-Version 164.08.00
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
type GetIPInterfaceNumberOfEntriesResponse struct {
	XMLName                    xml.Name // rather for debug purpose
	IPInterfaceNumberOfEntries int      `xml:"NewIPInterfaceNumberOfEntries"` // default=0
}
