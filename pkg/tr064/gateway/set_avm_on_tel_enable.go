package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetAvmOnTelEnable AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'SetEnable', Fritz!Box-System-Version 164.08.00
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func SetAvmOnTelEnable(session *soap.SoapSession, enable bool) (tr064model.SetAvmOnTelEnableResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("SetEnable").
		AddBoolParam("NewEnable", enable).
		Do()
	if err != nil {
		return tr064model.SetAvmOnTelEnableResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetAvmOnTelEnableResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetAvmOnTelEnableResponse{}, err
	}
	return result, nil
}
