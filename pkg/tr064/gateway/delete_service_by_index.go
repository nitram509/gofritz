package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// DeleteServiceByIndex AUTO-GENERATED (do not edit) code from [x_myfritzSCPD],
// based on SOAP action 'DeleteServiceByIndex', Fritz!Box-System-Version 164.08.00
//
// [x_myfritzSCPD]: http://fritz.box:49000/x_myfritzSCPD.xml
func DeleteServiceByIndex(session *soap.SoapSession, index int) (tr064model.DeleteServiceByIndexResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_myfritz").
		Uri("urn:dslforum-org:service:X_AVM-DE_MyFritz:1").
		Action("DeleteServiceByIndex").
		AddIntParam("NewIndex", index).
		Do()
	if err != nil {
		return tr064model.DeleteServiceByIndexResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.DeleteServiceByIndexResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.DeleteServiceByIndexResponse{}, err
	}
	return result, nil
}
