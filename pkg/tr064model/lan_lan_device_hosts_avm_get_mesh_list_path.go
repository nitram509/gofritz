package tr064model

import "encoding/xml"

// XavmGetMeshListPathResponse AUTO-GENERATED (do not edit) model from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetMeshListPath', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
type XavmGetMeshListPathResponse struct {
	XMLName      xml.Name // rather for debug purpose
	MeshListPath string   `xml:"NewX_AVM-DE_MeshListPath"` //
}
