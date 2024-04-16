package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// AddUSPController AUTO-GENERATED (do not edit) code from [x_uspcontrollerSCPD],
// based on SOAP action 'AddUSPController', Fritz!Box-System-Version 164.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
func AddUSPController(session *soap.SoapSession, enable bool, endpointId string, mTP string, hostname string, path string, port int, useTLS bool, accessRightSmarthome bool, accessRightMesh bool, accessRightInternet bool, accessRightSystem bool, accessRightController bool, accessRightWiFi bool, accessRightVoip bool, username string, password string) (tr064model.AddUSPControllerResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_uspcontroller").
		Uri("urn:dslforum-org:service:X_AVM-DE_USPController:1").
		Action("AddUSPController").
		AddBoolParam("NewEnable", enable).
		AddStringParam("NewEndpointID", endpointId).
		AddStringParam("NewMTP", mTP).
		AddStringParam("NewHostname", hostname).
		AddStringParam("NewPath", path).
		AddIntParam("NewPort", port).
		AddBoolParam("NewUseTLS", useTLS).
		AddBoolParam("NewAccessRightSmarthome", accessRightSmarthome).
		AddBoolParam("NewAccessRightMesh", accessRightMesh).
		AddBoolParam("NewAccessRightInternet", accessRightInternet).
		AddBoolParam("NewAccessRightSystem", accessRightSystem).
		AddBoolParam("NewAccessRightController", accessRightController).
		AddBoolParam("NewAccessRightWiFi", accessRightWiFi).
		AddBoolParam("NewAccessRightVoIP", accessRightVoip).
		AddStringParam("NewUsername", username).
		AddStringParam("NewPassword", password).
		Do().Body.Data
	result := tr064model.AddUSPControllerResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
