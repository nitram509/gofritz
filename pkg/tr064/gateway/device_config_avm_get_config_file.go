package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmGetConfigFile AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetConfigFile', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func XavmGetConfigFile(session *soap.SoapSession) (tr064model.XavmGetConfigFileResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("X_AVM-DE_GetConfigFile").
		Do().Body.Data
	result := tr064model.XavmGetConfigFileResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
