package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetVoIPEnableAreaCode AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'SetVoIPEnableAreaCode', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func SetVoIPEnableAreaCode(session *soap.SoapSession, voipAccountIndex int, voipEnableAreaCode bool) (tr064model.SetVoIPEnableAreaCodeResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("SetVoIPEnableAreaCode").
		AddIntParam("NewVoIPAccountIndex", voipAccountIndex).
		AddBoolParam("NewVoIPEnableAreaCode", voipEnableAreaCode).
		Do()
	if err != nil {
		return tr064model.SetVoIPEnableAreaCodeResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetVoIPEnableAreaCodeResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetVoIPEnableAreaCodeResponse{}, err
	}
	return result, nil
}
