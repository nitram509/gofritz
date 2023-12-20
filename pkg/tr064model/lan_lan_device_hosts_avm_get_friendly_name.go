package tr064model

import "encoding/xml"

// XavmGetFriendlyNameResponse AUTO-GENERATED (do not edit) model from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetFriendlyName', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
type XavmGetFriendlyNameResponse struct {
	XMLName      xml.Name // rather for debug purpose
	FriendlyName string   `xml:"NewX_AVM-DE_FriendlyName"` //
}
