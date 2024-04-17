package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1XavmSetWPSEnable AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetWPSEnable', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1XavmSetWPSEnable(session *soap.SoapSession, avmWpsEnable bool) (tr064model.XavmSetWPSEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("X_AVM-DE_SetWPSEnable").
		AddBoolParam("NewX_AVM-DE_WPSEnable", avmWpsEnable).
		Do().Body.Data
	result := tr064model.XavmSetWPSEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
