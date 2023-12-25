package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan1XavmSetStickSurfEnable AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetStickSurfEnable', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1XavmSetStickSurfEnable(session *soap.SoapSession, stickSurfEnable bool) (tr064model.XavmSetStickSurfEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("X_AVM-DE_SetStickSurfEnable").
		AddBoolParam("NewStickSurfEnable", stickSurfEnable).
		Do().Body.Data
	result := tr064model.XavmSetStickSurfEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
