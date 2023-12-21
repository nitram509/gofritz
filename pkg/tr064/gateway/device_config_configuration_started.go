package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// ConfigurationStarted AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'ConfigurationStarted', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func ConfigurationStarted(session *soap.SoapSession) (tr064model.ConfigurationStartedResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("ConfigurationStarted").
		Do().Body.Data
	result := tr064model.ConfigurationStartedResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
