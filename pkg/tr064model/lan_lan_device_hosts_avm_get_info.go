package tr064model

import "encoding/xml"

// XavmGetLanDeviceHostsInfoResponse AUTO-GENERATED (do not edit) model from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetInfo', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
type XavmGetLanDeviceHostsInfoResponse struct {
	XMLName              xml.Name // rather for debug purpose
	FriendlynameMinChars int      `xml:"NewX_AVM-DE_FriendlynameMinChars"` // default=1
	FriendlynameMaxChars int      `xml:"NewX_AVM-DE_FriendlynameMaxChars"` // default=64
	HostnameMinChars     int      `xml:"NewX_AVM-DE_HostnameMinChars"`     // default=0
	HostnameMaxChars     int      `xml:"NewX_AVM-DE_HostnameMaxChars"`     // default=63
	HostnameAllowedChars string   `xml:"NewX_AVM-DE_HostnameAllowedChars"` // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-
}
