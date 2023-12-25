package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmAddVoIPAccount AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_AddVoIPAccount', Fritz!Box-System-Version 141.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmAddVoIPAccount(session *soap.SoapSession, voipAccountIndex int, voipRegistrar string, voipNumber string, voipUsername string, voipPassword string, voipOutboundProxy string, voipSTUNServer string) (tr064model.XavmAddVoIPAccountResponse, error) {
	bodyData := soap.NewSoapRequest(session).
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
		Do().Body.Data
	result := tr064model.XavmAddVoIPAccountResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
