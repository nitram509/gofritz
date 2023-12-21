package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetConnectionRequestAuthentication AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'SetConnectionRequestAuthentication', Fritz!Box-System-Version 164.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func SetConnectionRequestAuthentication(session *soap.SoapSession) (tr064model.SetConnectionRequestAuthenticationResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("SetConnectionRequestAuthentication").
		Do().Body.Data
	result := tr064model.SetConnectionRequestAuthenticationResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
