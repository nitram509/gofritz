package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// FactoryReset AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'FactoryReset', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func FactoryReset(session *soap.SoapSession) (tr064model.FactoryResetResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("FactoryReset").
		Do().Body.Data
	result := tr064model.FactoryResetResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
