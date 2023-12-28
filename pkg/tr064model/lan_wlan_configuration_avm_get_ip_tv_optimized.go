package tr064model

import "encoding/xml"

// XavmGetIPTVOptimizedResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetIPTVOptimized', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type XavmGetIPTVOptimizedResponse struct {
	XMLName      xml.Name // rather for debug purpose
	IPTVoptimize bool     `xml:"NewX_AVM-DE_IPTVoptimize"` // default=0
}
