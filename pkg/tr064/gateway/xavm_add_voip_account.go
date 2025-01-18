package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmAddVoIPAccount AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_AddVoIPAccount', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmAddVoIPAccount(session *soap.SoapSession, voipAccountIndex int, voipRegistrar string, voipNumber string, voipUsername string, voipPassword string, voipOutboundProxy string, voipSTUNServer string) (tr064model.XavmAddVoIPAccountResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_AddVoIPAccount").
		AddIntParam("NewVoIPAccountIndex", voipAccountIndex).
		AddStringParam("NewVoIPRegistrar", voipRegistrar).
		AddStringParam("NewVoIPNumber", voipNumber).
		AddStringParam("NewVoIPUsername", voipUsername).
		AddStringParam("NewVoIPPassword", voipPassword).
		AddStringParam("NewVoIPOutboundProxy", voipOutboundProxy).
		AddStringParam("NewVoIPSTUNServer", voipSTUNServer).
		Do()
	if err != nil {
		return tr064model.XavmAddVoIPAccountResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmAddVoIPAccountResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmAddVoIPAccountResponse{}, err
	}
	return result, nil
}
