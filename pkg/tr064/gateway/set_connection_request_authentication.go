package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetConnectionRequestAuthentication AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'SetConnectionRequestAuthentication', Fritz!Box-System-Version 164.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func SetConnectionRequestAuthentication(session *soap.SoapSession, connectionRequestUsername string, connectionRequestPassword string) (tr064model.SetConnectionRequestAuthenticationResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("SetConnectionRequestAuthentication").
		AddStringParam("NewConnectionRequestUsername", connectionRequestUsername).
		AddStringParam("NewConnectionRequestPassword", connectionRequestPassword).
		Do().Body.Data
	result := tr064model.SetConnectionRequestAuthenticationResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
