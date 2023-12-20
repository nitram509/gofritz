package tr064model

import "encoding/xml"

// XavmGetWLANDeviceListPathResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetWLANDeviceListPath', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type XavmGetWLANDeviceListPathResponse struct {
	XMLName            xml.Name // rather for debug purpose
	WLANDeviceListPath string   `xml:"NewX_AVM-DE_WLANDeviceListPath"` //
}
