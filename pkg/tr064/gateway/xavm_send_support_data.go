package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSendSupportData AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_SendSupportData', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func XavmSendSupportData(session *soap.SoapSession, avmSupportDataMode string) (tr064model.XavmSendSupportDataResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("X_AVM-DE_SendSupportData").
		AddStringParam("NewX_AVM-DE_SupportDataMode", avmSupportDataMode).
		Do()
	if err != nil {
		return tr064model.XavmSendSupportDataResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSendSupportDataResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSendSupportDataResponse{}, err
	}
	return result, nil
}
