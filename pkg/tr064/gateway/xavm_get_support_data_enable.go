package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetSupportDataEnable AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetSupportDataEnable', Fritz!Box-System-Version 164.08.00
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func XavmGetSupportDataEnable(session *soap.SoapSession) (tr064model.XavmGetSupportDataEnableResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("X_AVM-DE_GetSupportDataEnable").
		Do()
	if err != nil {
		return tr064model.XavmGetSupportDataEnableResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetSupportDataEnableResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetSupportDataEnableResponse{}, err
	}
	return result, nil
}
