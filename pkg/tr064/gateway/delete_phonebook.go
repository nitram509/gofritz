package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// DeletePhonebook AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'DeletePhonebook', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func DeletePhonebook(session *soap.SoapSession, phonebookId int, phonebookExtraId string) (tr064model.DeletePhonebookResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("DeletePhonebook").
		AddIntParam("NewPhonebookID", phonebookId).
		AddStringParam("NewPhonebookExtraID", phonebookExtraId).
		Do().Body.Data
	result := tr064model.DeletePhonebookResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
