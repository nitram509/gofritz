package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// AddUSPController AUTO-GENERATED (do not edit) code from [x_uspcontrollerSCPD],
// based on SOAP action 'AddUSPController', Fritz!Box-System-Version 164.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
func AddUSPController(session *soap.SoapSession) (tr064model.AddUSPControllerResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_uspcontroller").
		Uri("urn:dslforum-org:service:X_AVM-DE_USPController:1").
		Action("AddUSPController").
		Do().Body.Data
	result := tr064model.AddUSPControllerResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
