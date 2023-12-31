package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan1SetWlanConfigurationEnable AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetEnable', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1SetWlanConfigurationEnable(session *soap.SoapSession, enable bool) (tr064model.SetWlanConfigurationEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("SetEnable").
		AddBoolParam("NewEnable", enable).
		Do().Body.Data
	result := tr064model.SetWlanConfigurationEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
