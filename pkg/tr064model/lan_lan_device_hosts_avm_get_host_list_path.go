package tr064model

import "encoding/xml"

// XavmGetHostListPathResponse AUTO-GENERATED (do not edit) model from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetHostListPath', Fritz!Box-System-Version 141.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
type XavmGetHostListPathResponse struct {
	XMLName      xml.Name // rather for debug purpose
	HostListPath string   `xml:"NewX_AVM-DE_HostListPath"` //
}
