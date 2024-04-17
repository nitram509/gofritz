package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// DeleteMessage AUTO-GENERATED (do not edit) code from [x_tamSCPD],
// based on SOAP action 'DeleteMessage', Fritz!Box-System-Version 164.07.57
//
// [x_tamSCPD]: http://fritz.box:49000/x_tamSCPD.xml
func DeleteMessage(session *soap.SoapSession, index int, messageIndex int) (tr064model.DeleteMessageResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_tam").
		Uri("urn:dslforum-org:service:X_AVM-DE_TAM:1").
		Action("DeleteMessage").
		AddIntParam("NewIndex", index).
		AddIntParam("NewMessageIndex", messageIndex).
		Do()
	if err != nil {
		return tr064model.DeleteMessageResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.DeleteMessageResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.DeleteMessageResponse{}, err
	}
	return result, nil
}
