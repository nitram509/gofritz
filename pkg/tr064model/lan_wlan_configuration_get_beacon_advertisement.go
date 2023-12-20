package tr064model

// GetBeaconAdvertisementResponse auto generated model from [wlanconfigSCPD],
// based on SOAP action 'GetBeaconAdvertisement', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetBeaconAdvertisementResponse struct {
	BeaconAdvertisementEnabled bool `xml:"NewBeaconAdvertisementEnabled"` // default=0
}
