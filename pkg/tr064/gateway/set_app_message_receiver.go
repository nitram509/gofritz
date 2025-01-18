package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetAppMessageReceiver AUTO-GENERATED (do not edit) code from [x_appsetupSCPD],
// based on SOAP action 'SetAppMessageReceiver', Fritz!Box-System-Version 164.08.00
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
func SetAppMessageReceiver(session *soap.SoapSession, appId string, cryptAlgos string, appAvmAddress string, appAvmPasswordHash string) (tr064model.SetAppMessageReceiverResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_appsetup").
		Uri("urn:dslforum-org:service:X_AVM-DE_AppSetup:1").
		Action("SetAppMessageReceiver").
		AddStringParam("NewAppId", appId).
		AddStringParam("NewCryptAlgos", cryptAlgos).
		AddStringParam("NewAppAVMAddress", appAvmAddress).
		AddStringParam("NewAppAVMPasswordHash", appAvmPasswordHash).
		Do()
	if err != nil {
		return tr064model.SetAppMessageReceiverResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetAppMessageReceiverResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetAppMessageReceiverResponse{}, err
	}
	return result, nil
}
