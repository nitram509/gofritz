package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetVoIPStatus AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetVoIPStatus', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmGetVoIPStatus(session *soap.SoapSession, voipAccountIndex int) (tr064model.XavmGetVoIPStatusResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_GetVoIPStatus").
		AddIntParam("NewVoIPAccountIndex", voipAccountIndex).
		Do().Body.Data
	result := tr064model.XavmGetVoIPStatusResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
