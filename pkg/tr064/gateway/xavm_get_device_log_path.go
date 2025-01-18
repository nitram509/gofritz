package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetDeviceLogPath AUTO-GENERATED (do not edit) code from [deviceinfoSCPD],
// based on SOAP action 'X_AVM-DE_GetDeviceLogPath', Fritz!Box-System-Version 164.08.00
//
// [deviceinfoSCPD]: http://fritz.box:49000/deviceinfoSCPD.xml
func XavmGetDeviceLogPath(session *soap.SoapSession) (tr064model.XavmGetDeviceLogPathResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceinfo").
		Uri("urn:dslforum-org:service:DeviceInfo:1").
		Action("X_AVM-DE_GetDeviceLogPath").
		Do()
	if err != nil {
		return tr064model.XavmGetDeviceLogPathResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetDeviceLogPathResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetDeviceLogPathResponse{}, err
	}
	return result, nil
}
