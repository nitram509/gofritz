package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Reboot AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'Reboot', Fritz!Box-System-Version 164.08.00
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func Reboot(session *soap.SoapSession) (tr064model.RebootResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("Reboot").
		Do()
	if err != nil {
		return tr064model.RebootResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.RebootResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.RebootResponse{}, err
	}
	return result, nil
}
