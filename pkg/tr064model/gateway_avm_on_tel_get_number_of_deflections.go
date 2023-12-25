package tr064model

import "encoding/xml"

// GetNumberOfDeflectionsResponse AUTO-GENERATED (do not edit) model from [x_contactSCPD],
// based on SOAP action 'GetNumberOfDeflections', Fritz!Box-System-Version 141.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetNumberOfDeflectionsResponse struct {
	XMLName             xml.Name // rather for debug purpose
	NumberOfDeflections int      `xml:"NewNumberOfDeflections"` // default=0
}
