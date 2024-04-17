package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// ConfigurationFinished AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'ConfigurationFinished', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func ConfigurationFinished(session *soap.SoapSession) (tr064model.ConfigurationFinishedResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("ConfigurationFinished").
		Do().Body.Data
	result := tr064model.ConfigurationFinishedResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
