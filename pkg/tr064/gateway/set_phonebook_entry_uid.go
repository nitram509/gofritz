package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetPhonebookEntryUID AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'SetPhonebookEntryUID', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func SetPhonebookEntryUID(session *soap.SoapSession, phonebookId int, phonebookEntryData string) (tr064model.SetPhonebookEntryUIDResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("SetPhonebookEntryUID").
		AddIntParam("NewPhonebookID", phonebookId).
		AddStringParam("NewPhonebookEntryData", phonebookEntryData).
		Do()
	if err != nil {
		return tr064model.SetPhonebookEntryUIDResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetPhonebookEntryUIDResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetPhonebookEntryUIDResponse{}, err
	}
	return result, nil
}
