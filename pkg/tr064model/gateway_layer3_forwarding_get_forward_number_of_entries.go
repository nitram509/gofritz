package tr064model

import "encoding/xml"

// GetForwardNumberOfEntriesResponse AUTO-GENERATED (do not edit) model from [layer3forwardingSCPD],
// based on SOAP action 'GetForwardNumberOfEntries', Fritz!Box-System-Version 164.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
type GetForwardNumberOfEntriesResponse struct {
	XMLName                xml.Name // rather for debug purpose
	ForwardNumberOfEntries int      `xml:"NewForwardNumberOfEntries"` // default=0
}
