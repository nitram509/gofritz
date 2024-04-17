package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetVoIPCommonAreaCode AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'SetVoIPCommonAreaCode', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func SetVoIPCommonAreaCode(session *soap.SoapSession, voipAreaCode string) (tr064model.SetVoIPCommonAreaCodeResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("SetVoIPCommonAreaCode").
		AddStringParam("NewVoIPAreaCode", voipAreaCode).
		Do()
	if err != nil {
		return tr064model.SetVoIPCommonAreaCodeResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetVoIPCommonAreaCodeResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetVoIPCommonAreaCodeResponse{}, err
	}
	return result, nil
}
