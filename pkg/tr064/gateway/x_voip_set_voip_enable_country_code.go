package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetVoIPEnableCountryCode AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'SetVoIPEnableCountryCode', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func SetVoIPEnableCountryCode(session *soap.SoapSession, voipAccountIndex int, voipEnableCountryCode bool) (tr064model.SetVoIPEnableCountryCodeResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("SetVoIPEnableCountryCode").
		AddIntParam("NewVoIPAccountIndex", voipAccountIndex).
		AddBoolParam("NewVoIPEnableCountryCode", voipEnableCountryCode).
		Do().Body.Data
	result := tr064model.SetVoIPEnableCountryCodeResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
