package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan3SetSecurityKeys AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetSecurityKeys', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan3SetSecurityKeys(session *soap.SoapSession, wepKey0 string, wepKey1 string, wepKey2 string, wepKey3 string, preSharedKey string, keyPassphrase string) (tr064model.SetSecurityKeysResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("SetSecurityKeys").
		AddStringParam("NewWEPKey0", wepKey0).
		AddStringParam("NewWEPKey1", wepKey1).
		AddStringParam("NewWEPKey2", wepKey2).
		AddStringParam("NewWEPKey3", wepKey3).
		AddStringParam("NewPreSharedKey", preSharedKey).
		AddStringParam("NewKeyPassphrase", keyPassphrase).
		Do().Body.Data
	result := tr064model.SetSecurityKeysResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
