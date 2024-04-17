package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetVoIPCommonAreaCode AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_SetVoIPCommonAreaCode', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmSetVoIPCommonAreaCode(session *soap.SoapSession, avmOKZ string, avmOKZPrefix string) (tr064model.XavmSetVoIPCommonAreaCodeResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_SetVoIPCommonAreaCode").
		AddStringParam("NewX_AVM-DE_OKZ", avmOKZ).
		AddStringParam("NewX_AVM-DE_OKZPrefix", avmOKZPrefix).
		Do()
	if err != nil {
		return tr064model.XavmSetVoIPCommonAreaCodeResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetVoIPCommonAreaCodeResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetVoIPCommonAreaCodeResponse{}, err
	}
	return result, nil
}
