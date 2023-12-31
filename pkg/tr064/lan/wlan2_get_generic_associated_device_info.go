package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2GetGenericAssociatedDeviceInfo AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetGenericAssociatedDeviceInfo', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2GetGenericAssociatedDeviceInfo(session *soap.SoapSession, associatedDeviceIndex int) (tr064model.GetGenericAssociatedDeviceInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("GetGenericAssociatedDeviceInfo").
		AddIntParam("NewAssociatedDeviceIndex", associatedDeviceIndex).
		Do().Body.Data
	result := tr064model.GetGenericAssociatedDeviceInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
