package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetAvmTamEnable AUTO-GENERATED (do not edit) code from [x_tamSCPD],
// based on SOAP action 'SetEnable', Fritz!Box-System-Version 164.07.57
//
// [x_tamSCPD]: http://fritz.box:49000/x_tamSCPD.xml
func SetAvmTamEnable(session *soap.SoapSession, index int, enable bool) (tr064model.SetAvmTamEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_tam").
		Uri("urn:dslforum-org:service:X_AVM-DE_TAM:1").
		Action("SetEnable").
		AddIntParam("NewIndex", index).
		AddBoolParam("NewEnable", enable).
		Do().Body.Data
	result := tr064model.SetAvmTamEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
