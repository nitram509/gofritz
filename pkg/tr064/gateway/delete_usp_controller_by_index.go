package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// DeleteUSPControllerByIndex AUTO-GENERATED (do not edit) code from [x_uspcontrollerSCPD],
// based on SOAP action 'DeleteUSPControllerByIndex', Fritz!Box-System-Version 164.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
func DeleteUSPControllerByIndex(session *soap.SoapSession, index int) (tr064model.DeleteUSPControllerByIndexResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_uspcontroller").
		Uri("urn:dslforum-org:service:X_AVM-DE_USPController:1").
		Action("DeleteUSPControllerByIndex").
		AddIntParam("NewIndex", index).
		Do().Body.Data
	result := tr064model.DeleteUSPControllerByIndexResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
