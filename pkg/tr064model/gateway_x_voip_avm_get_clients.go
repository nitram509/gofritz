package tr064model

import "encoding/xml"

// XavmGetClientsResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetClients', Fritz!Box-System-Version 141.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetClientsResponse struct {
	XMLName    xml.Name // rather for debug purpose
	ClientList string   `xml:"NewX_AVM-DE_ClientList"` //
}
