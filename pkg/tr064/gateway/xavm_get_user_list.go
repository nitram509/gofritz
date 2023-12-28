package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmGetUserList AUTO-GENERATED (do not edit) code from [lanconfigsecuritySCPD],
// based on SOAP action 'X_AVM-DE_GetUserList', Fritz!Box-System-Version 164.07.57
//
// [lanconfigsecuritySCPD]: http://fritz.box:49000/lanconfigsecuritySCPD.xml
func XavmGetUserList(session *soap.SoapSession) (tr064model.XavmGetUserListResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanconfigsecurity").
		Uri("urn:dslforum-org:service:LANConfigSecurity:1").
		Action("X_AVM-DE_GetUserList").
		Do().Body.Data
	result := tr064model.XavmGetUserListResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
