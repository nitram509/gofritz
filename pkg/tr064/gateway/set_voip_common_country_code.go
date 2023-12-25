package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetVoIPCommonCountryCode AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'SetVoIPCommonCountryCode', Fritz!Box-System-Version 141.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func SetVoIPCommonCountryCode(session *soap.SoapSession, voipCountryCode string) (tr064model.SetVoIPCommonCountryCodeResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("SetVoIPCommonCountryCode").
		AddStringParam("NewVoIPCountryCode", voipCountryCode).
		Do().Body.Data
	result := tr064model.SetVoIPCommonCountryCodeResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
