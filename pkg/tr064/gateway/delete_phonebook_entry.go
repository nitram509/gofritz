package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// DeletePhonebookEntry AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'DeletePhonebookEntry', Fritz!Box-System-Version 164.08.00
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func DeletePhonebookEntry(session *soap.SoapSession, phonebookId int, phonebookEntryId int) (tr064model.DeletePhonebookEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("DeletePhonebookEntry").
		AddIntParam("NewPhonebookID", phonebookId).
		AddIntParam("NewPhonebookEntryID", phonebookEntryId).
		Do()
	if err != nil {
		return tr064model.DeletePhonebookEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.DeletePhonebookEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.DeletePhonebookEntryResponse{}, err
	}
	return result, nil
}
