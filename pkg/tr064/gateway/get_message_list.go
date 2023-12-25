package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetMessageList AUTO-GENERATED (do not edit) code from [x_tamSCPD],
// based on SOAP action 'GetMessageList', Fritz!Box-System-Version 141.07.57
//
// [x_tamSCPD]: http://fritz.box:49000/x_tamSCPD.xml
func GetMessageList(session *soap.SoapSession, index int) (tr064model.GetMessageListResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_tam").
		Uri("urn:dslforum-org:service:X_AVM-DE_TAM:1").
		Action("GetMessageList").
		AddIntParam("NewIndex", index).
		Do().Body.Data
	result := tr064model.GetMessageListResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
