package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan3GetSpecificAssociatedDeviceInfo AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetSpecificAssociatedDeviceInfo', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan3GetSpecificAssociatedDeviceInfo(session *soap.SoapSession, associatedDeviceMacAddress string) (tr064model.GetSpecificAssociatedDeviceInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("GetSpecificAssociatedDeviceInfo").
		AddStringParam("NewAssociatedDeviceMACAddress", associatedDeviceMacAddress).
		Do()
	if err != nil {
		return tr064model.GetSpecificAssociatedDeviceInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetSpecificAssociatedDeviceInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetSpecificAssociatedDeviceInfoResponse{}, err
	}
	return result, nil
}
