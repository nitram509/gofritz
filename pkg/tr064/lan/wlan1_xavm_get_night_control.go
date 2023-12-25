package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan1XavmGetNightControl AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetNightControl', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1XavmGetNightControl(session *soap.SoapSession) (tr064model.XavmGetNightControlResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("X_AVM-DE_GetNightControl").
		Do().Body.Data
	result := tr064model.XavmGetNightControlResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
