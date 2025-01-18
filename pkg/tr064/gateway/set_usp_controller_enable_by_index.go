package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetUSPControllerEnableByIndex AUTO-GENERATED (do not edit) code from [x_uspcontrollerSCPD],
// based on SOAP action 'SetUSPControllerEnableByIndex', Fritz!Box-System-Version 164.08.00
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
func SetUSPControllerEnableByIndex(session *soap.SoapSession, index int, enable bool) (tr064model.SetUSPControllerEnableByIndexResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_uspcontroller").
		Uri("urn:dslforum-org:service:X_AVM-DE_USPController:1").
		Action("SetUSPControllerEnableByIndex").
		AddIntParam("NewIndex", index).
		AddBoolParam("NewEnable", enable).
		Do()
	if err != nil {
		return tr064model.SetUSPControllerEnableByIndexResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetUSPControllerEnableByIndexResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetUSPControllerEnableByIndexResponse{}, err
	}
	return result, nil
}
