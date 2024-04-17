package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetClient AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_SetClient', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmSetClient(session *soap.SoapSession, avmClientIndex int, avmClientPassword string, avmPhoneName string, avmOutGoingNumber string) (tr064model.XavmSetClientResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_SetClient").
		AddIntParam("NewX_AVM-DE_ClientIndex", avmClientIndex).
		AddStringParam("NewX_AVM-DE_ClientPassword", avmClientPassword).
		AddStringParam("NewX_AVM-DE_PhoneName", avmPhoneName).
		AddStringParam("NewX_AVM-DE_OutGoingNumber", avmOutGoingNumber).
		Do()
	if err != nil {
		return tr064model.XavmSetClientResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetClientResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetClientResponse{}, err
	}
	return result, nil
}
