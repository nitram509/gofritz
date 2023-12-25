package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan1GetSecurityKeys AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetSecurityKeys', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1GetSecurityKeys(session *soap.SoapSession) (tr064model.GetSecurityKeysResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("GetSecurityKeys").
		Do().Body.Data
	result := tr064model.GetSecurityKeysResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
