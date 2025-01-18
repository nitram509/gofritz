package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan2SetWlanConfigurationEnable AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetEnable', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2SetWlanConfigurationEnable(session *soap.SoapSession, enable bool) (tr064model.SetWlanConfigurationEnableResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("SetEnable").
		AddBoolParam("NewEnable", enable).
		Do()
	if err != nil {
		return tr064model.SetWlanConfigurationEnableResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetWlanConfigurationEnableResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetWlanConfigurationEnableResponse{}, err
	}
	return result, nil
}
