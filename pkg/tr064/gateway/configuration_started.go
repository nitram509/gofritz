package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// ConfigurationStarted AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'ConfigurationStarted', Fritz!Box-System-Version 164.08.00
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func ConfigurationStarted(session *soap.SoapSession, sessionId string) (tr064model.ConfigurationStartedResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("ConfigurationStarted").
		AddStringParam("NewSessionID", sessionId).
		Do()
	if err != nil {
		return tr064model.ConfigurationStartedResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.ConfigurationStartedResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.ConfigurationStartedResponse{}, err
	}
	return result, nil
}
