package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetClient4 AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_SetClient4', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmSetClient4(session *soap.SoapSession, avmClientIndex int, avmClientPassword string, avmClientUsername string, avmPhoneName string, avmClientId string, avmOutGoingNumber string, avmInComingNumbers string) (tr064model.XavmSetClient4Response, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_SetClient4").
		AddIntParam("NewX_AVM-DE_ClientIndex", avmClientIndex).
		AddStringParam("NewX_AVM-DE_ClientPassword", avmClientPassword).
		AddStringParam("NewX_AVM-DE_ClientUsername", avmClientUsername).
		AddStringParam("NewX_AVM-DE_PhoneName", avmPhoneName).
		AddStringParam("NewX_AVM-DE_ClientId", avmClientId).
		AddStringParam("NewX_AVM-DE_OutGoingNumber", avmOutGoingNumber).
		AddStringParam("NewX_AVM-DE_InComingNumbers", avmInComingNumbers).
		Do()
	if err != nil {
		return tr064model.XavmSetClient4Response{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetClient4Response{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetClient4Response{}, err
	}
	return result, nil
}
