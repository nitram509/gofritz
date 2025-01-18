package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetDeviceInfo AUTO-GENERATED (do not edit) code from [deviceinfoSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [deviceinfoSCPD]: http://fritz.box:49000/deviceinfoSCPD.xml
func GetDeviceInfo(session *soap.SoapSession) (tr064model.GetDeviceInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceinfo").
		Uri("urn:dslforum-org:service:DeviceInfo:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetDeviceInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetDeviceInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetDeviceInfoResponse{}, err
	}
	return result, nil
}
