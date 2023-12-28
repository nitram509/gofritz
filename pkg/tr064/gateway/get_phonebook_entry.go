package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetPhonebookEntry AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'GetPhonebookEntry', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func GetPhonebookEntry(session *soap.SoapSession, phonebookId int, phonebookEntryId int) (tr064model.GetPhonebookEntryResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("GetPhonebookEntry").
		AddIntParam("NewPhonebookID", phonebookId).
		AddIntParam("NewPhonebookEntryID", phonebookEntryId).
		Do().Body.Data
	result := tr064model.GetPhonebookEntryResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
