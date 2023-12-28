package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan1XavmGetIPTVOptimized AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetIPTVOptimized', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1XavmGetIPTVOptimized(session *soap.SoapSession) (tr064model.XavmGetIPTVOptimizedResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("X_AVM-DE_GetIPTVOptimized").
		Do().Body.Data
	result := tr064model.XavmGetIPTVOptimizedResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
