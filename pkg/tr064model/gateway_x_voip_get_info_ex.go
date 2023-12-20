package tr064model

import "encoding/xml"

// GetInfoExResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'GetInfoEx', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type GetInfoExResponse struct {
	XMLName                    xml.Name // rather for debug purpose
	VoIPNumberMinChars         int      `xml:"NewVoIPNumberMinChars"`                  // default=1
	VoIPNumberMaxChars         int      `xml:"NewVoIPNumberMaxChars"`                  // default=128
	VoIPNumberAllowedChars     string   `xml:"NewVoIPNumberAllowedChars"`              // default=0123456789+
	VoIPUsernameMinChars       int      `xml:"NewVoIPUsernameMinChars"`                // default=4
	VoIPUsernameMaxChars       int      `xml:"NewVoIPUsernameMaxChars"`                // default=128
	VoIPUsernameAllowedChars   string   `xml:"NewVoIPUsernameAllowedChars"`            // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-._@
	VoIPPasswordMinChars       int      `xml:"NewVoIPPasswordMinChars"`                // default=3
	VoIPPasswordMaxChars       int      `xml:"NewVoIPPasswordMaxChars"`                // default=64
	VoIPPasswordAllowedChars   string   `xml:"NewVoIPPasswordAllowedChars"`            // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-._
	VoIPRegistrarMinChars      int      `xml:"NewVoIPRegistrarMinChars"`               // default=1
	VoIPRegistrarMaxChars      int      `xml:"NewVoIPRegistrarMaxChars"`               // default=64
	VoIPRegistrarAllowedChars  string   `xml:"NewVoIPRegistrarAllowedChars"`           // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-._
	VoIPSTUNServerMinChars     int      `xml:"NewVoIPSTUNServerMinChars"`              // default=0
	VoIPSTUNServerMaxChars     int      `xml:"NewVoIPSTUNServerMaxChars"`              // default=255
	VoIPSTUNServerAllowedChars string   `xml:"NewVoIPSTUNServerAllowedChars"`          // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-._:
	ClientUsernameMinChars     int      `xml:"NewX_AVM-DE_ClientUsernameMinChars"`     // default=4
	ClientUsernameMaxChars     int      `xml:"NewX_AVM-DE_ClientUsernameMaxChars"`     // default=64
	ClientUsernameAllowedChars string   `xml:"NewX_AVM-DE_ClientUsernameAllowedChars"` // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-._
	ClientPasswordMinChars     int      `xml:"NewX_AVM-DE_ClientPasswordMinChars"`     // default=8
	ClientPasswordMaxChars     int      `xml:"NewX_AVM-DE_ClientPasswordMaxChars"`     // default=64
	ClientPasswordAllowedChars string   `xml:"NewX_AVM-DE_ClientPasswordAllowedChars"` // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-._
}
