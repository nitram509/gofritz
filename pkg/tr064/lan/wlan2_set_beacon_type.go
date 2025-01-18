package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan2SetBeaconType AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetBeaconType', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2SetBeaconType(session *soap.SoapSession, beaconType string) (tr064model.SetBeaconTypeResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("SetBeaconType").
		AddStringParam("NewBeaconType", beaconType).
		Do()
	if err != nil {
		return tr064model.SetBeaconTypeResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetBeaconTypeResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetBeaconTypeResponse{}, err
	}
	return result, nil
}
