package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmSetClient3 AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_SetClient3', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmSetClient3(session *soap.SoapSession, avmClientIndex int, avmClientPassword string, avmClientId string, avmPhoneName string, avmOutGoingNumber string, avmInComingNumbers string, avmExternalRegistration bool) (tr064model.XavmSetClient3Response, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_SetClient3").
		AddIntParam("NewX_AVM-DE_ClientIndex", avmClientIndex).
		AddStringParam("NewX_AVM-DE_ClientPassword", avmClientPassword).
		AddStringParam("NewX_AVM-DE_ClientId", avmClientId).
		AddStringParam("NewX_AVM-DE_PhoneName", avmPhoneName).
		AddStringParam("NewX_AVM-DE_OutGoingNumber", avmOutGoingNumber).
		AddStringParam("NewX_AVM-DE_InComingNumbers", avmInComingNumbers).
		AddBoolParam("NewX_AVM-DE_ExternalRegistration", avmExternalRegistration).
		Do().Body.Data
	result := tr064model.XavmSetClient3Response{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
