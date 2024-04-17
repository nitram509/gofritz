package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetCallBarringEntry AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'GetCallBarringEntry', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func GetCallBarringEntry(session *soap.SoapSession, phonebookEntryId int) (tr064model.GetCallBarringEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("GetCallBarringEntry").
		AddIntParam("NewPhonebookEntryID", phonebookEntryId).
		Do()
	if err != nil {
		return tr064model.GetCallBarringEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetCallBarringEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetCallBarringEntryResponse{}, err
	}
	return result, nil
}
