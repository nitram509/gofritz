package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetSupportDataInfo AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetSupportDataInfo', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func XavmGetSupportDataInfo(session *soap.SoapSession) (tr064model.XavmGetSupportDataInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("X_AVM-DE_GetSupportDataInfo").
		Do().Body.Data
	result := tr064model.XavmGetSupportDataInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
