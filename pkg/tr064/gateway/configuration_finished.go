package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// ConfigurationFinished AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'ConfigurationFinished', Fritz!Box-System-Version 164.08.00
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func ConfigurationFinished(session *soap.SoapSession) (tr064model.ConfigurationFinishedResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("ConfigurationFinished").
		Do()
	if err != nil {
		return tr064model.ConfigurationFinishedResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.ConfigurationFinishedResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.ConfigurationFinishedResponse{}, err
	}
	return result, nil
}
