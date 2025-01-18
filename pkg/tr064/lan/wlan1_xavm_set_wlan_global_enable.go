package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1XavmSetWLANGlobalEnable AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetWLANGlobalEnable', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1XavmSetWLANGlobalEnable(session *soap.SoapSession, avmWlanGlobalEnable bool) (tr064model.XavmSetWLANGlobalEnableResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("X_AVM-DE_SetWLANGlobalEnable").
		AddBoolParam("NewX_AVM-DE_WLANGlobalEnable", avmWlanGlobalEnable).
		Do()
	if err != nil {
		return tr064model.XavmSetWLANGlobalEnableResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetWLANGlobalEnableResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetWLANGlobalEnableResponse{}, err
	}
	return result, nil
}
