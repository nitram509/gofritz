package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// DeleteCallBarringEntryUID AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'DeleteCallBarringEntryUID', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func DeleteCallBarringEntryUID(session *soap.SoapSession, phonebookEntryUniqueId int) (tr064model.DeleteCallBarringEntryUIDResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("DeleteCallBarringEntryUID").
		AddIntParam("NewPhonebookEntryUniqueID", phonebookEntryUniqueId).
		Do().Body.Data
	result := tr064model.DeleteCallBarringEntryUIDResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
