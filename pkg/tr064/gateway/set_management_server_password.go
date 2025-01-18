package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetManagementServerPassword AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'SetManagementServerPassword', Fritz!Box-System-Version 164.08.00
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func SetManagementServerPassword(session *soap.SoapSession, password string) (tr064model.SetManagementServerPasswordResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("SetManagementServerPassword").
		AddStringParam("NewPassword", password).
		Do()
	if err != nil {
		return tr064model.SetManagementServerPasswordResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetManagementServerPasswordResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetManagementServerPasswordResponse{}, err
	}
	return result, nil
}
