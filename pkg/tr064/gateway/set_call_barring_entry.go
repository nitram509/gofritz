package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetCallBarringEntry AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'SetCallBarringEntry', Fritz!Box-System-Version 164.08.00
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func SetCallBarringEntry(session *soap.SoapSession, phonebookEntryData string) (tr064model.SetCallBarringEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("SetCallBarringEntry").
		AddStringParam("NewPhonebookEntryData", phonebookEntryData).
		Do()
	if err != nil {
		return tr064model.SetCallBarringEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetCallBarringEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetCallBarringEntryResponse{}, err
	}
	return result, nil
}
