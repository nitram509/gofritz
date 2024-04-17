package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetDeviceLog AUTO-GENERATED (do not edit) code from [deviceinfoSCPD],
// based on SOAP action 'GetDeviceLog', Fritz!Box-System-Version 164.07.57
//
// [deviceinfoSCPD]: http://fritz.box:49000/deviceinfoSCPD.xml
func GetDeviceLog(session *soap.SoapSession) (tr064model.GetDeviceLogResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceinfo").
		Uri("urn:dslforum-org:service:DeviceInfo:1").
		Action("GetDeviceLog").
		Do()
	if err != nil {
		return tr064model.GetDeviceLogResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetDeviceLogResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetDeviceLogResponse{}, err
	}
	return result, nil
}
