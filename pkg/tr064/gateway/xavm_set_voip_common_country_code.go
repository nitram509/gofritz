package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetVoIPCommonCountryCode AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_SetVoIPCommonCountryCode', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmSetVoIPCommonCountryCode(session *soap.SoapSession, avmLKZ string, avmLKZPrefix string) (tr064model.XavmSetVoIPCommonCountryCodeResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_SetVoIPCommonCountryCode").
		AddStringParam("NewX_AVM-DE_LKZ", avmLKZ).
		AddStringParam("NewX_AVM-DE_LKZPrefix", avmLKZPrefix).
		Do()
	if err != nil {
		return tr064model.XavmSetVoIPCommonCountryCodeResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetVoIPCommonCountryCodeResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetVoIPCommonCountryCodeResponse{}, err
	}
	return result, nil
}
