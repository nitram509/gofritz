package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetXVoipConfig AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func SetXVoipConfig(session *soap.SoapSession, faxT38Enable bool, voiceCoding string) (tr064model.SetXVoipConfigResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("SetConfig").
		AddBoolParam("NewFaxT38Enable", faxT38Enable).
		AddStringParam("NewVoiceCoding", voiceCoding).
		Do()
	if err != nil {
		return tr064model.SetXVoipConfigResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetXVoipConfigResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetXVoipConfigResponse{}, err
	}
	return result, nil
}
