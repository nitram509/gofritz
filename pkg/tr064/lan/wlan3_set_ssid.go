package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan3SetSSID AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetSSID', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan3SetSSID(session *soap.SoapSession, ssid string) (tr064model.SetSSIDResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("SetSSID").
		AddStringParam("NewSSID", ssid).
		Do()
	if err != nil {
		return tr064model.SetSSIDResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetSSIDResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetSSIDResponse{}, err
	}
	return result, nil
}
