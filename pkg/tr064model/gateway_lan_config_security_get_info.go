package tr064model

// GetLanConfigSecurityInfoResponse AUTO-GENERATED (do not edit) model from [lanconfigsecuritySCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [lanconfigsecuritySCPD]: http://fritz.box:49000/lanconfigsecuritySCPD.xml
type GetLanConfigSecurityInfoResponse struct {
	MaxCharsPassword        int    `xml:"NewMaxCharsPassword"`                 // default=32
	MinCharsPassword        int    `xml:"NewMinCharsPassword"`                 // default=0
	AllowedCharsPassword    string `xml:"NewAllowedCharsPassword"`             // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~
	AllowedCharsUsername    string `xml:"NewAllowedCharsUsername"`             // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz -._
	IsDefaultPasswordActive bool   `xml:"NewX_AVM-DE_IsDefaultPasswordActive"` // default=0
}
