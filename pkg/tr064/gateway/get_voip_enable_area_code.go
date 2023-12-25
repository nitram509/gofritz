package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetVoIPEnableAreaCode AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'GetVoIPEnableAreaCode', Fritz!Box-System-Version 141.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func GetVoIPEnableAreaCode(session *soap.SoapSession, voipAccountIndex int) (tr064model.GetVoIPEnableAreaCodeResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("GetVoIPEnableAreaCode").
		AddIntParam("NewVoIPAccountIndex", voipAccountIndex).
		Do().Body.Data
	result := tr064model.GetVoIPEnableAreaCodeResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
