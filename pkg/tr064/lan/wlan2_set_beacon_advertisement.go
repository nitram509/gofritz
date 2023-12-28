package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2SetBeaconAdvertisement AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetBeaconAdvertisement', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2SetBeaconAdvertisement(session *soap.SoapSession, beaconAdvertisementEnabled bool) (tr064model.SetBeaconAdvertisementResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("SetBeaconAdvertisement").
		AddBoolParam("NewBeaconAdvertisementEnabled", beaconAdvertisementEnabled).
		Do().Body.Data
	result := tr064model.SetBeaconAdvertisementResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
