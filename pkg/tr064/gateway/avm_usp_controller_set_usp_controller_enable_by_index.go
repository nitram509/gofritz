package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetUSPControllerEnableByIndex AUTO-GENERATED (do not edit) code from [x_uspcontrollerSCPD],
// based on SOAP action 'SetUSPControllerEnableByIndex', Fritz!Box-System-Version 164.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
func SetUSPControllerEnableByIndex(session *soap.SoapSession) (tr064model.SetUSPControllerEnableByIndexResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_uspcontroller").
		Uri("urn:dslforum-org:service:X_AVM-DE_USPController:1").
		Action("SetUSPControllerEnableByIndex").
		Do().Body.Data
	result := tr064model.SetUSPControllerEnableByIndexResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
