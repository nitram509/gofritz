package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan3GetGenericAssociatedDeviceInfo AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetGenericAssociatedDeviceInfo', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan3GetGenericAssociatedDeviceInfo(session *soap.SoapSession, associatedDeviceIndex int) (tr064model.GetGenericAssociatedDeviceInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("GetGenericAssociatedDeviceInfo").
		AddIntParam("NewAssociatedDeviceIndex", associatedDeviceIndex).
		Do()
	if err != nil {
		return tr064model.GetGenericAssociatedDeviceInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetGenericAssociatedDeviceInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetGenericAssociatedDeviceInfoResponse{}, err
	}
	return result, nil
}
