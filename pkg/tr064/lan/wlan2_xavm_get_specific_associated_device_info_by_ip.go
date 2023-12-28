package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2XavmGetSpecificAssociatedDeviceInfoByIp AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetSpecificAssociatedDeviceInfoByIp', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2XavmGetSpecificAssociatedDeviceInfoByIp(session *soap.SoapSession, associatedDeviceIpAddress string) (tr064model.XavmGetSpecificAssociatedDeviceInfoByIpResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("X_AVM-DE_GetSpecificAssociatedDeviceInfoByIp").
		AddStringParam("NewAssociatedDeviceIPAddress", associatedDeviceIpAddress).
		Do().Body.Data
	result := tr064model.XavmGetSpecificAssociatedDeviceInfoByIpResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
