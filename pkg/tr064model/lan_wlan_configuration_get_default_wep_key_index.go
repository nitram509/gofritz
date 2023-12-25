package tr064model

import "encoding/xml"

// GetDefaultWEPKeyIndexResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetDefaultWEPKeyIndex', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetDefaultWEPKeyIndexResponse struct {
	XMLName            xml.Name // rather for debug purpose
	DefaultWEPKeyIndex int      `xml:"NewDefaultWEPKeyIndex"` // default=0
}
