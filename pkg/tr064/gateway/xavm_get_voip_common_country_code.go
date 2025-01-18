package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetVoIPCommonCountryCode AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetVoIPCommonCountryCode', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmGetVoIPCommonCountryCode(session *soap.SoapSession) (tr064model.XavmGetVoIPCommonCountryCodeResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_GetVoIPCommonCountryCode").
		Do()
	if err != nil {
		return tr064model.XavmGetVoIPCommonCountryCodeResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetVoIPCommonCountryCodeResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetVoIPCommonCountryCodeResponse{}, err
	}
	return result, nil
}
