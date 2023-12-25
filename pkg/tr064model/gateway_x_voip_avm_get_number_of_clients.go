package tr064model

import "encoding/xml"

// XavmGetNumberOfClientsResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetNumberOfClients', Fritz!Box-System-Version 141.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetNumberOfClientsResponse struct {
	XMLName         xml.Name // rather for debug purpose
	NumberOfClients int      `xml:"NewX_AVM-DE_NumberOfClients"` // default=0
}
