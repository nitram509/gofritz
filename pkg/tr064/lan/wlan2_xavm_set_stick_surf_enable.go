package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan2XavmSetStickSurfEnable AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetStickSurfEnable', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2XavmSetStickSurfEnable(session *soap.SoapSession, stickSurfEnable bool) (tr064model.XavmSetStickSurfEnableResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("X_AVM-DE_SetStickSurfEnable").
		AddBoolParam("NewStickSurfEnable", stickSurfEnable).
		Do()
	if err != nil {
		return tr064model.XavmSetStickSurfEnableResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetStickSurfEnableResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetStickSurfEnableResponse{}, err
	}
	return result, nil
}
