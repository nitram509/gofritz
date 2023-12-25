package tr064model

import "encoding/xml"

// XavmGetWPSInfoResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetWPSInfo', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type XavmGetWPSInfoResponse struct {
	XMLName   xml.Name // rather for debug purpose
	WPSMode   string   `xml:"NewX_AVM-DE_WPSMode"`   // oneOf=[other,stop,pbc]
	WPSStatus string   `xml:"NewX_AVM-DE_WPSStatus"` // oneOf=[off,inactive,active,success,err_common,err_timeout,err_reconfig,err_internal,err_abort,]
}
