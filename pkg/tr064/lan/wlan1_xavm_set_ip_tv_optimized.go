package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1XavmSetIPTVOptimized AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetIPTVOptimized', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1XavmSetIPTVOptimized(session *soap.SoapSession, avmIpTvoptimize bool) (tr064model.XavmSetIPTVOptimizedResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("X_AVM-DE_SetIPTVOptimized").
		AddBoolParam("NewX_AVM-DE_IPTVoptimize", avmIpTvoptimize).
		Do()
	if err != nil {
		return tr064model.XavmSetIPTVOptimizedResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetIPTVOptimizedResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetIPTVOptimizedResponse{}, err
	}
	return result, nil
}
