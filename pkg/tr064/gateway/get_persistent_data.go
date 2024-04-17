package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetPersistentData AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'GetPersistentData', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func GetPersistentData(session *soap.SoapSession) (tr064model.GetPersistentDataResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("GetPersistentData").
		Do().Body.Data
	result := tr064model.GetPersistentDataResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
