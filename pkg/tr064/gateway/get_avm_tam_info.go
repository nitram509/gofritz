package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetAvmTamInfo AUTO-GENERATED (do not edit) code from [x_tamSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_tamSCPD]: http://fritz.box:49000/x_tamSCPD.xml
func GetAvmTamInfo(session *soap.SoapSession, index int) (tr064model.GetAvmTamInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_tam").
		Uri("urn:dslforum-org:service:X_AVM-DE_TAM:1").
		Action("GetInfo").
		AddIntParam("NewIndex", index).
		Do().Body.Data
	result := tr064model.GetAvmTamInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
