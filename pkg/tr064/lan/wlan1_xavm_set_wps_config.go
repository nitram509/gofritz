package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1XavmSetWPSConfig AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetWPSConfig', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1XavmSetWPSConfig(session *soap.SoapSession, avmWpsMode string) (tr064model.XavmSetWPSConfigResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("X_AVM-DE_SetWPSConfig").
		AddStringParam("NewX_AVM-DE_WPSMode", avmWpsMode).
		Do()
	if err != nil {
		return tr064model.XavmSetWPSConfigResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetWPSConfigResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetWPSConfigResponse{}, err
	}
	return result, nil
}
