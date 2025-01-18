package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1XavmGetSpecificAssociatedDeviceInfoByIp AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetSpecificAssociatedDeviceInfoByIp', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1XavmGetSpecificAssociatedDeviceInfoByIp(session *soap.SoapSession, associatedDeviceIpAddress string) (tr064model.XavmGetSpecificAssociatedDeviceInfoByIpResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("X_AVM-DE_GetSpecificAssociatedDeviceInfoByIp").
		AddStringParam("NewAssociatedDeviceIPAddress", associatedDeviceIpAddress).
		Do()
	if err != nil {
		return tr064model.XavmGetSpecificAssociatedDeviceInfoByIpResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetSpecificAssociatedDeviceInfoByIpResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetSpecificAssociatedDeviceInfoByIpResponse{}, err
	}
	return result, nil
}
