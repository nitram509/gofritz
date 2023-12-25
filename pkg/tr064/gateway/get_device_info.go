package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetDeviceInfo AUTO-GENERATED (do not edit) code from [deviceinfoSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 141.07.57
//
// [deviceinfoSCPD]: http://fritz.box:49000/deviceinfoSCPD.xml
func GetDeviceInfo(session *soap.SoapSession) (tr064model.GetDeviceInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceinfo").
		Uri("urn:dslforum-org:service:DeviceInfo:1").
		Action("GetInfo").
		Do().Body.Data
	result := tr064model.GetDeviceInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
