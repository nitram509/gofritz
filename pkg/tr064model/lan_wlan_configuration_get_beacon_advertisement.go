package tr064model

import "encoding/xml"

// GetBeaconAdvertisementResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetBeaconAdvertisement', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetBeaconAdvertisementResponse struct {
	XMLName                    xml.Name // rather for debug purpose
	BeaconAdvertisementEnabled bool     `xml:"NewBeaconAdvertisementEnabled"` // default=0
}
