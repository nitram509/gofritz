package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetVoIPCommonAreaCode AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'GetVoIPCommonAreaCode', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func GetVoIPCommonAreaCode(session *soap.SoapSession) (tr064model.GetVoIPCommonAreaCodeResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("GetVoIPCommonAreaCode").
		Do()
	if err != nil {
		return tr064model.GetVoIPCommonAreaCodeResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetVoIPCommonAreaCodeResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetVoIPCommonAreaCodeResponse{}, err
	}
	return result, nil
}
