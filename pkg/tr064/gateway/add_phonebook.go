package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// AddPhonebook AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'AddPhonebook', Fritz!Box-System-Version 164.08.00
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func AddPhonebook(session *soap.SoapSession, phonebookExtraId string, phonebookName string) (tr064model.AddPhonebookResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("AddPhonebook").
		AddStringParam("NewPhonebookExtraID", phonebookExtraId).
		AddStringParam("NewPhonebookName", phonebookName).
		Do()
	if err != nil {
		return tr064model.AddPhonebookResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.AddPhonebookResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.AddPhonebookResponse{}, err
	}
	return result, nil
}
