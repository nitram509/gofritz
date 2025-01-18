package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetExistingVoIPNumbers AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'GetExistingVoIPNumbers', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func GetExistingVoIPNumbers(session *soap.SoapSession) (tr064model.GetExistingVoIPNumbersResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("GetExistingVoIPNumbers").
		Do()
	if err != nil {
		return tr064model.GetExistingVoIPNumbersResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetExistingVoIPNumbersResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetExistingVoIPNumbersResponse{}, err
	}
	return result, nil
}
