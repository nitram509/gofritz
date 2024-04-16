package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// DeletePhonebookEntryUID AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'DeletePhonebookEntryUID', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func DeletePhonebookEntryUID(session *soap.SoapSession, phonebookId int, phonebookEntryUniqueId int) (tr064model.DeletePhonebookEntryUIDResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("DeletePhonebookEntryUID").
		AddIntParam("NewPhonebookID", phonebookId).
		AddIntParam("NewPhonebookEntryUniqueID", phonebookEntryUniqueId).
		Do().Body.Data
	result := tr064model.DeletePhonebookEntryUIDResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
