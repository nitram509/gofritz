package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetManagementServerUsername AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'SetManagementServerUsername', Fritz!Box-System-Version 164.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func SetManagementServerUsername(session *soap.SoapSession, username string) (tr064model.SetManagementServerUsernameResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("SetManagementServerUsername").
		AddStringParam("NewUsername", username).
		Do().Body.Data
	result := tr064model.SetManagementServerUsernameResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
