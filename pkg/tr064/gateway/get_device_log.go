package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetDeviceLog AUTO-GENERATED (do not edit) code from [deviceinfoSCPD],
// based on SOAP action 'GetDeviceLog', Fritz!Box-System-Version 141.07.57
//
// [deviceinfoSCPD]: http://fritz.box:49000/deviceinfoSCPD.xml
func GetDeviceLog(session *soap.SoapSession) (tr064model.GetDeviceLogResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceinfo").
		Uri("urn:dslforum-org:service:DeviceInfo:1").
		Action("GetDeviceLog").
		Do().Body.Data
	result := tr064model.GetDeviceLogResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
