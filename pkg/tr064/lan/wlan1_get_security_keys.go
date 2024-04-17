package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1GetSecurityKeys AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetSecurityKeys', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1GetSecurityKeys(session *soap.SoapSession) (tr064model.GetSecurityKeysResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("GetSecurityKeys").
		Do()
	if err != nil {
		return tr064model.GetSecurityKeysResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetSecurityKeysResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetSecurityKeysResponse{}, err
	}
	return result, nil
}
