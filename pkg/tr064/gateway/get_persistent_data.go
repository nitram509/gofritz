package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetPersistentData AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'GetPersistentData', Fritz!Box-System-Version 164.08.00
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func GetPersistentData(session *soap.SoapSession) (tr064model.GetPersistentDataResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("GetPersistentData").
		Do()
	if err != nil {
		return tr064model.GetPersistentDataResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetPersistentDataResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetPersistentDataResponse{}, err
	}
	return result, nil
}
