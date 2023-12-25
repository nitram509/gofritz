package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// MarkMessage AUTO-GENERATED (do not edit) code from [x_tamSCPD],
// based on SOAP action 'MarkMessage', Fritz!Box-System-Version 141.07.57
//
// [x_tamSCPD]: http://fritz.box:49000/x_tamSCPD.xml
func MarkMessage(session *soap.SoapSession, index int, messageIndex int, markedAsRead bool) (tr064model.MarkMessageResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_tam").
		Uri("urn:dslforum-org:service:X_AVM-DE_TAM:1").
		Action("MarkMessage").
		AddIntParam("NewIndex", index).
		AddIntParam("NewMessageIndex", messageIndex).
		AddBoolParam("NewMarkedAsRead", markedAsRead).
		Do().Body.Data
	result := tr064model.MarkMessageResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
