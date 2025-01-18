package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1XavmGetIPTVOptimized AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetIPTVOptimized', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1XavmGetIPTVOptimized(session *soap.SoapSession) (tr064model.XavmGetIPTVOptimizedResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("X_AVM-DE_GetIPTVOptimized").
		Do()
	if err != nil {
		return tr064model.XavmGetIPTVOptimizedResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetIPTVOptimizedResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetIPTVOptimizedResponse{}, err
	}
	return result, nil
}
