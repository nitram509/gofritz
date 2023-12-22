package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetAvmOnTelConfig AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func SetAvmOnTelConfig(session *soap.SoapSession, enable bool, url string, serviceId string, username string, password string, name string) (tr064model.SetAvmOnTelConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("SetConfig").
		AddBoolParam("NewEnable", enable).
		AddStringParam("NewUrl", url).
		AddStringParam("NewServiceId", serviceId).
		AddStringParam("NewUsername", username).
		AddStringParam("NewPassword", password).
		AddStringParam("NewName", name).
		Do().Body.Data
	result := tr064model.SetAvmOnTelConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
