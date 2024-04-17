package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// X_GenerateUUID AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'X_GenerateUUID', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func X_GenerateUUID(session *soap.SoapSession) (tr064model.X_GenerateUUIDResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("X_GenerateUUID").
		Do().Body.Data
	result := tr064model.X_GenerateUUIDResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
