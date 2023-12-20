package tr064model

import "encoding/xml"

// XavmGetUserListResponse AUTO-GENERATED (do not edit) model from [lanconfigsecuritySCPD],
// based on SOAP action 'X_AVM-DE_GetUserList', Fritz!Box-System-Version 164.07.57
//
// [lanconfigsecuritySCPD]: http://fritz.box:49000/lanconfigsecuritySCPD.xml
type XavmGetUserListResponse struct {
	XMLName  xml.Name // rather for debug purpose
	UserList string   `xml:"NewX_AVM-DE_UserList"` //
}
