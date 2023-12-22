package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetUSPControllerByIndex AUTO-GENERATED (do not edit) code from [x_uspcontrollerSCPD],
// based on SOAP action 'GetUSPControllerByIndex', Fritz!Box-System-Version 164.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
func GetUSPControllerByIndex(session *soap.SoapSession, index int) (tr064model.GetUSPControllerByIndexResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_uspcontroller").
		Uri("urn:dslforum-org:service:X_AVM-DE_USPController:1").
		Action("GetUSPControllerByIndex").
		AddIntParam("NewIndex", index).
		Do().Body.Data
	result := tr064model.GetUSPControllerByIndexResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
