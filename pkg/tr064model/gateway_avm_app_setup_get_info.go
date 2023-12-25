package tr064model

import "encoding/xml"

// GetAvmAppSetupInfoResponse AUTO-GENERATED (do not edit) model from [x_appsetupSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 141.07.57
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
type GetAvmAppSetupInfoResponse struct {
	XMLName                        xml.Name // rather for debug purpose
	MinCharsAppId                  int      `xml:"NewMinCharsAppId"`                  // default=1
	MaxCharsAppId                  int      `xml:"NewMaxCharsAppId"`                  // default=256
	AllowedCharsAppId              string   `xml:"NewAllowedCharsAppId"`              // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	MinCharsAppDisplayName         int      `xml:"NewMinCharsAppDisplayName"`         // default=1
	MaxCharsAppDisplayName         int      `xml:"NewMaxCharsAppDisplayName"`         // default=256
	MinCharsAppUsername            int      `xml:"NewMinCharsAppUsername"`            // default=1
	MaxCharsAppUsername            int      `xml:"NewMaxCharsAppUsername"`            // default=32
	AllowedCharsAppUsername        string   `xml:"NewAllowedCharsAppUsername"`        // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	MinCharsAppPassword            int      `xml:"NewMinCharsAppPassword"`            // default=8
	MaxCharsAppPassword            int      `xml:"NewMaxCharsAppPassword"`            // default=32
	AllowedCharsAppPassword        string   `xml:"NewAllowedCharsAppPassword"`        // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~
	MinCharsIPSecIdentifier        int      `xml:"NewMinCharsIPSecIdentifier"`        // default=1
	MaxCharsIPSecIdentifier        int      `xml:"NewMaxCharsIPSecIdentifier"`        // default=256
	AllowedCharsIPSecIdentifier    string   `xml:"NewAllowedCharsIPSecIdentifier"`    // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	AllowedCharsCryptAlgos         string   `xml:"NewAllowedCharsCryptAlgos"`         // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_.
	AllowedCharsAppAVMAddress      string   `xml:"NewAllowedCharsAppAVMAddress"`      // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_.
	MinCharsFilter                 int      `xml:"NewMinCharsFilter"`                 // default=0
	MaxCharsFilter                 int      `xml:"NewMaxCharsFilter"`                 // default=1024
	AllowedCharsFilter             string   `xml:"NewAllowedCharsFilter"`             // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz,+:-_
	MinCharsIPSecPreSharedKey      int      `xml:"NewMinCharsIPSecPreSharedKey"`      // default=1
	MaxCharsIPSecPreSharedKey      int      `xml:"NewMaxCharsIPSecPreSharedKey"`      // default=64
	AllowedCharsIPSecPreSharedKey  string   `xml:"NewAllowedCharsIPSecPreSharedKey"`  // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	MinCharsIPSecXauthUsername     int      `xml:"NewMinCharsIPSecXauthUsername"`     // default=1
	MaxCharsIPSecXauthUsername     int      `xml:"NewMaxCharsIPSecXauthUsername"`     // default=256
	AllowedCharsIPSecXauthUsername string   `xml:"NewAllowedCharsIPSecXauthUsername"` // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	MinCharsIPSecXauthPassword     int      `xml:"NewMinCharsIPSecXauthPassword"`     // default=1
	MaxCharsIPSecXauthPassword     int      `xml:"NewMaxCharsIPSecXauthPassword"`     // default=128
	AllowedCharsIPSecXauthPassword string   `xml:"NewAllowedCharsIPSecXauthPassword"` // default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
}
