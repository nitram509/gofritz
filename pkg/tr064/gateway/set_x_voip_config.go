package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetXVoipConfig AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 141.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func SetXVoipConfig(session *soap.SoapSession, faxT38Enable bool, voiceCoding string) (tr064model.SetXVoipConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("SetConfig").
		AddBoolParam("NewFaxT38Enable", faxT38Enable).
		AddStringParam("NewVoiceCoding", voiceCoding).
		Do().Body.Data
	result := tr064model.SetXVoipConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
