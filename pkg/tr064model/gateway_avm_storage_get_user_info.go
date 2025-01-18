package tr064model

import "encoding/xml"

// GetUserInfoResponse AUTO-GENERATED (do not edit) model from [x_storageSCPD],
// based on SOAP action 'GetUserInfo', Fritz!Box-System-Version 164.08.00
//
// [x_storageSCPD]: http://fritz.box:49000/x_storageSCPD.xml
type GetUserInfoResponse struct {
	XMLName               xml.Name // rather for debug purpose
	Enable                bool     `xml:"NewEnable"`                         // default=0
	Username              string   `xml:"NewUsername"`                       //
	NetworkAccessReadOnly bool     `xml:"NewX_AVM-DE_NetworkAccessReadOnly"` // default=0
}
