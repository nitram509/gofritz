package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetPhonebookEntry AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'SetPhonebookEntry', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func SetPhonebookEntry(session *soap.SoapSession, phonebookId int, phonebookEntryId int, phonebookEntryData string) (tr064model.SetPhonebookEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("SetPhonebookEntry").
		AddIntParam("NewPhonebookID", phonebookId).
		AddIntParam("NewPhonebookEntryID", phonebookEntryId).
		AddStringParam("NewPhonebookEntryData", phonebookEntryData).
		Do()
	if err != nil {
		return tr064model.SetPhonebookEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetPhonebookEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetPhonebookEntryResponse{}, err
	}
	return result, nil
}
