package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetInfoByIndex AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'GetInfoByIndex', Fritz!Box-System-Version 164.08.00
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func GetInfoByIndex(session *soap.SoapSession, index int) (tr064model.GetInfoByIndexResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("GetInfoByIndex").
		AddIntParam("NewIndex", index).
		Do()
	if err != nil {
		return tr064model.GetInfoByIndexResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetInfoByIndexResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetInfoByIndexResponse{}, err
	}
	return result, nil
}
