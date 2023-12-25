package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2GetSpecificAssociatedDeviceInfo AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetSpecificAssociatedDeviceInfo', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2GetSpecificAssociatedDeviceInfo(session *soap.SoapSession, associatedDeviceMacAddress string) (tr064model.GetSpecificAssociatedDeviceInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("GetSpecificAssociatedDeviceInfo").
		AddStringParam("NewAssociatedDeviceMACAddress", associatedDeviceMacAddress).
		Do().Body.Data
	result := tr064model.GetSpecificAssociatedDeviceInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
