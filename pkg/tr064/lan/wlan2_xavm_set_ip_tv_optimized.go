package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2XavmSetIPTVOptimized AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetIPTVOptimized', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2XavmSetIPTVOptimized(session *soap.SoapSession, avmIpTvoptimize bool) (tr064model.XavmSetIPTVOptimizedResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("X_AVM-DE_SetIPTVOptimized").
		AddBoolParam("NewX_AVM-DE_IPTVoptimize", avmIpTvoptimize).
		Do().Body.Data
	result := tr064model.XavmSetIPTVOptimizedResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
