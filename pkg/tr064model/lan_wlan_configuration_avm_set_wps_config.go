package tr064model

import "encoding/xml"

// XavmSetWPSConfigResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetWPSConfig', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type XavmSetWPSConfigResponse struct {
	XMLName   xml.Name // rather for debug purpose
	WPSStatus string   `xml:"NewX_AVM-DE_WPSStatus"` // oneOf=[off,inactive,active,success,err_common,err_timeout,err_reconfig,err_internal,err_abort,]
}
