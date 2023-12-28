package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan1GetWlanConfigurationInfo AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1GetWlanConfigurationInfo(session *soap.SoapSession) (tr064model.GetWlanConfigurationInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("GetInfo").
		Do().Body.Data
	result := tr064model.GetWlanConfigurationInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
