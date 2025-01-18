package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetClient3 AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetClient3', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmGetClient3(session *soap.SoapSession, avmClientIndex int) (tr064model.XavmGetClient3Response, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_GetClient3").
		AddIntParam("NewX_AVM-DE_ClientIndex", avmClientIndex).
		Do()
	if err != nil {
		return tr064model.XavmGetClient3Response{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetClient3Response{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetClient3Response{}, err
	}
	return result, nil
}
