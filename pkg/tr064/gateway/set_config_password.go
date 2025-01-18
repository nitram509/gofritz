package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetConfigPassword AUTO-GENERATED (do not edit) code from [lanconfigsecuritySCPD],
// based on SOAP action 'SetConfigPassword', Fritz!Box-System-Version 164.08.00
//
// [lanconfigsecuritySCPD]: http://fritz.box:49000/lanconfigsecuritySCPD.xml
func SetConfigPassword(session *soap.SoapSession, password string) (tr064model.SetConfigPasswordResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanconfigsecurity").
		Uri("urn:dslforum-org:service:LANConfigSecurity:1").
		Action("SetConfigPassword").
		AddStringParam("NewPassword", password).
		Do()
	if err != nil {
		return tr064model.SetConfigPasswordResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetConfigPasswordResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetConfigPasswordResponse{}, err
	}
	return result, nil
}
