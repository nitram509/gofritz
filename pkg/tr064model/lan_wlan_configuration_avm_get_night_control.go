package tr064model

import "encoding/xml"

// XavmGetNightControlResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetNightControl', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type XavmGetNightControlResponse struct {
	XMLName                     xml.Name // rather for debug purpose
	NightControl                string   `xml:"NewNightControl"`                //
	NightTimeControlNoForcedOff bool     `xml:"NewNightTimeControlNoForcedOff"` // default=0
}
