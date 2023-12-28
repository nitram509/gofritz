package tr064model

import "encoding/xml"

// GetBeaconTypeResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetBeaconType', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetBeaconTypeResponse struct {
	XMLName             xml.Name // rather for debug purpose
	BeaconType          string   `xml:"NewBeaconType"`                   // oneOf=[none,Basic,WPA,11i,WPAand11i,WPA3,11iandWPA3,OWE,OWETrans]
	PossibleBeaconTypes string   `xml:"NewX_AVM-DE_PossibleBeaconTypes"` //
}
