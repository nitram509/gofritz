package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetManagementServerPassword AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'SetManagementServerPassword', Fritz!Box-System-Version 141.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func SetManagementServerPassword(session *soap.SoapSession, password string) (tr064model.SetManagementServerPasswordResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("SetManagementServerPassword").
		AddStringParam("NewPassword", password).
		Do().Body.Data
	result := tr064model.SetManagementServerPasswordResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
