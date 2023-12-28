package tr064model

import "encoding/xml"

// XavmCreateUrlSIDResponse AUTO-GENERATED (do not edit) model from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_CreateUrlSID', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
type XavmCreateUrlSIDResponse struct {
	XMLName xml.Name // rather for debug purpose
	UrlSID  string   `xml:"NewX_AVM-DE_UrlSID"` //
}
