package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan3XavmSetWPSEnable AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetWPSEnable', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan3XavmSetWPSEnable(session *soap.SoapSession, avmWpsEnable bool) (tr064model.XavmSetWPSEnableResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("X_AVM-DE_SetWPSEnable").
		AddBoolParam("NewX_AVM-DE_WPSEnable", avmWpsEnable).
		Do()
	if err != nil {
		return tr064model.XavmSetWPSEnableResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetWPSEnableResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetWPSEnableResponse{}, err
	}
	return result, nil
}
