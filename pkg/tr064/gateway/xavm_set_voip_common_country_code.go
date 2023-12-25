package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmSetVoIPCommonCountryCode AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_SetVoIPCommonCountryCode', Fritz!Box-System-Version 141.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmSetVoIPCommonCountryCode(session *soap.SoapSession, avmLKZ string, avmLKZPrefix string) (tr064model.XavmSetVoIPCommonCountryCodeResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_SetVoIPCommonCountryCode").
		AddStringParam("NewX_AVM-DE_LKZ", avmLKZ).
		AddStringParam("NewX_AVM-DE_LKZPrefix", avmLKZPrefix).
		Do().Body.Data
	result := tr064model.XavmSetVoIPCommonCountryCodeResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
