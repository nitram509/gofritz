package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetInfoByIndex AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'GetInfoByIndex', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func GetInfoByIndex(session *soap.SoapSession, index int) (tr064model.GetInfoByIndexResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("GetInfoByIndex").
		AddIntParam("NewIndex", index).
		Do().Body.Data
	result := tr064model.GetInfoByIndexResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
