package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// DeleteByIndex AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'DeleteByIndex', Fritz!Box-System-Version 141.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func DeleteByIndex(session *soap.SoapSession, index int) (tr064model.DeleteByIndexResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("DeleteByIndex").
		AddIntParam("NewIndex", index).
		Do().Body.Data
	result := tr064model.DeleteByIndexResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
