package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetUserList AUTO-GENERATED (do not edit) code from [lanconfigsecuritySCPD],
// based on SOAP action 'X_AVM-DE_GetUserList', Fritz!Box-System-Version 164.07.57
//
// [lanconfigsecuritySCPD]: http://fritz.box:49000/lanconfigsecuritySCPD.xml
func XavmGetUserList(session *soap.SoapSession) (tr064model.XavmGetUserListResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanconfigsecurity").
		Uri("urn:dslforum-org:service:LANConfigSecurity:1").
		Action("X_AVM-DE_GetUserList").
		Do()
	if err != nil {
		return tr064model.XavmGetUserListResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetUserListResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetUserListResponse{}, err
	}
	return result, nil
}
