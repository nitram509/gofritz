package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetPhonebook AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'GetPhonebook', Fritz!Box-System-Version 164.08.00
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func GetPhonebook(session *soap.SoapSession, phonebookId int) (tr064model.GetPhonebookResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("GetPhonebook").
		AddIntParam("NewPhonebookID", phonebookId).
		Do()
	if err != nil {
		return tr064model.GetPhonebookResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetPhonebookResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetPhonebookResponse{}, err
	}
	return result, nil
}
