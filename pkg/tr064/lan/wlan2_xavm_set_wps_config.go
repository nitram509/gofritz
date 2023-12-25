package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2XavmSetWPSConfig AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetWPSConfig', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2XavmSetWPSConfig(session *soap.SoapSession, avmWpsMode string) (tr064model.XavmSetWPSConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("X_AVM-DE_SetWPSConfig").
		AddStringParam("NewX_AVM-DE_WPSMode", avmWpsMode).
		Do().Body.Data
	result := tr064model.XavmSetWPSConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
