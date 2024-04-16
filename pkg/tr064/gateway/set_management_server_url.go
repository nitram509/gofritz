package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetManagementServerURL AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'SetManagementServerURL', Fritz!Box-System-Version 164.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func SetManagementServerURL(session *soap.SoapSession, url string) (tr064model.SetManagementServerURLResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("SetManagementServerURL").
		AddStringParam("NewURL", url).
		Do().Body.Data
	result := tr064model.SetManagementServerURLResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
