package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2SetSSID AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetSSID', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2SetSSID(session *soap.SoapSession, ssid string) (tr064model.SetSSIDResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("SetSSID").
		AddStringParam("NewSSID", ssid).
		Do().Body.Data
	result := tr064model.SetSSIDResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
