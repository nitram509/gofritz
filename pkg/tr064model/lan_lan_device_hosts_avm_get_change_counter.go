package tr064model

import "encoding/xml"

// XavmGetChangeCounterResponse AUTO-GENERATED (do not edit) model from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetChangeCounter', Fritz!Box-System-Version 141.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
type XavmGetChangeCounterResponse struct {
	XMLName       xml.Name // rather for debug purpose
	ChangeCounter int      `xml:"NewX_AVM-DE_ChangeCounter"` // default=0
}
