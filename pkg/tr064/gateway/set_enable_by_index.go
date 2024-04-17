package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetEnableByIndex AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'SetEnableByIndex', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func SetEnableByIndex(session *soap.SoapSession, index int, enable bool) (tr064model.SetEnableByIndexResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("SetEnableByIndex").
		AddIntParam("NewIndex", index).
		AddBoolParam("NewEnable", enable).
		Do().Body.Data
	result := tr064model.SetEnableByIndexResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
