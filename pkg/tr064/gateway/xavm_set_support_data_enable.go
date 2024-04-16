package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetSupportDataEnable AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetSupportDataEnable', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func XavmSetSupportDataEnable(session *soap.SoapSession, avmSupportDataEnabled bool) (tr064model.XavmSetSupportDataEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("X_AVM-DE_SetSupportDataEnable").
		AddBoolParam("NewX_AVM-DE_SupportDataEnabled", avmSupportDataEnabled).
		Do().Body.Data
	result := tr064model.XavmSetSupportDataEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
