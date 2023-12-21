package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetPersistentData AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'SetPersistentData', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func SetPersistentData(session *soap.SoapSession) (tr064model.SetPersistentDataResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("SetPersistentData").
		Do().Body.Data
	result := tr064model.SetPersistentDataResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
