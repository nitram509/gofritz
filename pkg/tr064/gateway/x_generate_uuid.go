package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// X_GenerateUUID AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'X_GenerateUUID', Fritz!Box-System-Version 164.08.00
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func X_GenerateUUID(session *soap.SoapSession) (tr064model.X_GenerateUUIDResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("X_GenerateUUID").
		Do()
	if err != nil {
		return tr064model.X_GenerateUUIDResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.X_GenerateUUIDResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.X_GenerateUUIDResponse{}, err
	}
	return result, nil
}
