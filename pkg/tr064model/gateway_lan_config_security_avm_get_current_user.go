package tr064model

import "encoding/xml"

// XavmGetCurrentUserResponse AUTO-GENERATED (do not edit) model from [lanconfigsecuritySCPD],
// based on SOAP action 'X_AVM-DE_GetCurrentUser', Fritz!Box-System-Version 164.07.57
//
// [lanconfigsecuritySCPD]: http://fritz.box:49000/lanconfigsecuritySCPD.xml
type XavmGetCurrentUserResponse struct {
	XMLName           xml.Name // rather for debug purpose
	CurrentUsername   string   `xml:"NewX_AVM-DE_CurrentUsername"`   //
	CurrentUserRights string   `xml:"NewX_AVM-DE_CurrentUserRights"` //
}
