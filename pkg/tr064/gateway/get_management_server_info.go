package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetManagementServerInfo AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func GetManagementServerInfo(session *soap.SoapSession) (tr064model.GetManagementServerInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("GetInfo").
		Do().Body.Data
	result := tr064model.GetManagementServerInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
