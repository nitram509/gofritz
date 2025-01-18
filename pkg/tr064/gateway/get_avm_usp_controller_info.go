package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAvmUspControllerInfo AUTO-GENERATED (do not edit) code from [x_uspcontrollerSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
func GetAvmUspControllerInfo(session *soap.SoapSession) (tr064model.GetAvmUspControllerInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_uspcontroller").
		Uri("urn:dslforum-org:service:X_AVM-DE_USPController:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetAvmUspControllerInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetAvmUspControllerInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetAvmUspControllerInfoResponse{}, err
	}
	return result, nil
}
