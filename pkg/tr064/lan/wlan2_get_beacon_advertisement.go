package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan2GetBeaconAdvertisement AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetBeaconAdvertisement', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2GetBeaconAdvertisement(session *soap.SoapSession) (tr064model.GetBeaconAdvertisementResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("GetBeaconAdvertisement").
		Do().Body.Data
	result := tr064model.GetBeaconAdvertisementResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
