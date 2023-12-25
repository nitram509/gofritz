package tr064model

import "encoding/xml"

// XavmGetActiveProviderResponse AUTO-GENERATED (do not edit) model from [wancommonifconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetActiveProvider', Fritz!Box-System-Version 141.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
type XavmGetActiveProviderResponse struct {
	XMLName  xml.Name // rather for debug purpose
	Provider string   `xml:"NewX_AVM-DE_Provider"` //
}
