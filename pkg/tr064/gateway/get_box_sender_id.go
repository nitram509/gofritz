package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetBoxSenderId AUTO-GENERATED (do not edit) code from [x_appsetupSCPD],
// based on SOAP action 'GetBoxSenderId', Fritz!Box-System-Version 164.08.00
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
func GetBoxSenderId(session *soap.SoapSession, appId string) (tr064model.GetBoxSenderIdResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_appsetup").
		Uri("urn:dslforum-org:service:X_AVM-DE_AppSetup:1").
		Action("GetBoxSenderId").
		AddStringParam("NewAppId", appId).
		Do()
	if err != nil {
		return tr064model.GetBoxSenderIdResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetBoxSenderIdResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetBoxSenderIdResponse{}, err
	}
	return result, nil
}
