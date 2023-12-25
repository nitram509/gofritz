package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan3XavmGetWPSInfo AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetWPSInfo', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan3XavmGetWPSInfo(session *soap.SoapSession) (tr064model.XavmGetWPSInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("X_AVM-DE_GetWPSInfo").
		Do().Body.Data
	result := tr064model.XavmGetWPSInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
