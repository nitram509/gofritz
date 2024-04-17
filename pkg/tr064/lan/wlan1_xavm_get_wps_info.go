package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1XavmGetWPSInfo AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetWPSInfo', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1XavmGetWPSInfo(session *soap.SoapSession) (tr064model.XavmGetWPSInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("X_AVM-DE_GetWPSInfo").
		Do().Body.Data
	result := tr064model.XavmGetWPSInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
