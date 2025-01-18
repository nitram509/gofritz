package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetManagementServerInfo AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func GetManagementServerInfo(session *soap.SoapSession) (tr064model.GetManagementServerInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetManagementServerInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetManagementServerInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetManagementServerInfoResponse{}, err
	}
	return result, nil
}
