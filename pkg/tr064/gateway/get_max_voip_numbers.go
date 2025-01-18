package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetMaxVoIPNumbers AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'GetMaxVoIPNumbers', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func GetMaxVoIPNumbers(session *soap.SoapSession) (tr064model.GetMaxVoIPNumbersResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("GetMaxVoIPNumbers").
		Do()
	if err != nil {
		return tr064model.GetMaxVoIPNumbersResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetMaxVoIPNumbersResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetMaxVoIPNumbersResponse{}, err
	}
	return result, nil
}
