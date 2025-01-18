package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// X_SetTR069Enable AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'X_SetTR069Enable', Fritz!Box-System-Version 164.08.00
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func X_SetTR069Enable(session *soap.SoapSession, tr069Enabled bool) (tr064model.X_SetTR069EnableResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("X_SetTR069Enable").
		AddBoolParam("NewTR069Enabled", tr069Enabled).
		Do()
	if err != nil {
		return tr064model.X_SetTR069EnableResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.X_SetTR069EnableResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.X_SetTR069EnableResponse{}, err
	}
	return result, nil
}
