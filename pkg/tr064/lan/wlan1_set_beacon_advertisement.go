package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1SetBeaconAdvertisement AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetBeaconAdvertisement', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1SetBeaconAdvertisement(session *soap.SoapSession, beaconAdvertisementEnabled bool) (tr064model.SetBeaconAdvertisementResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("SetBeaconAdvertisement").
		AddBoolParam("NewBeaconAdvertisementEnabled", beaconAdvertisementEnabled).
		Do()
	if err != nil {
		return tr064model.SetBeaconAdvertisementResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetBeaconAdvertisementResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetBeaconAdvertisementResponse{}, err
	}
	return result, nil
}
