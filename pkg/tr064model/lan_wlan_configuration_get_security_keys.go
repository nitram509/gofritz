package tr064model

import "encoding/xml"

// GetSecurityKeysResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetSecurityKeys', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetSecurityKeysResponse struct {
	XMLName       xml.Name // rather for debug purpose
	WEPKey0       string   `xml:"NewWEPKey0"`       //
	WEPKey1       string   `xml:"NewWEPKey1"`       //
	WEPKey2       string   `xml:"NewWEPKey2"`       //
	WEPKey3       string   `xml:"NewWEPKey3"`       //
	PreSharedKey  string   `xml:"NewPreSharedKey"`  //
	KeyPassphrase string   `xml:"NewKeyPassphrase"` //
}
