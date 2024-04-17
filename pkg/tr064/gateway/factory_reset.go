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
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("FactoryReset").
		Do()
	if err != nil {
		return tr064model.FactoryResetResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.FactoryResetResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.FactoryResetResponse{}, err
	}
	return result, nil
}
