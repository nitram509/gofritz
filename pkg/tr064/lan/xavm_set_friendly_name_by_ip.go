package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetFriendlyNameByIP AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_SetFriendlyNameByIP', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmSetFriendlyNameByIP(session *soap.SoapSession, ipAddress string, avmFriendlyName string) (tr064model.XavmSetFriendlyNameByIPResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_SetFriendlyNameByIP").
		AddStringParam("NewIPAddress", ipAddress).
		AddStringParam("NewX_AVM-DE_FriendlyName", avmFriendlyName).
		Do()
	if err != nil {
		return tr064model.XavmSetFriendlyNameByIPResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetFriendlyNameByIPResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetFriendlyNameByIPResponse{}, err
	}
	return result, nil
}
