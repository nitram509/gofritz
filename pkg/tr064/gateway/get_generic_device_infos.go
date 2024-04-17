package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetGenericDeviceInfos AUTO-GENERATED (do not edit) code from [x_homeautoSCPD],
// based on SOAP action 'GetGenericDeviceInfos', Fritz!Box-System-Version 164.07.57
//
// [x_homeautoSCPD]: http://fritz.box:49000/x_homeautoSCPD.xml
func GetGenericDeviceInfos(session *soap.SoapSession, index int) (tr064model.GetGenericDeviceInfosResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_homeauto").
		Uri("urn:dslforum-org:service:X_AVM-DE_Homeauto:1").
		Action("GetGenericDeviceInfos").
		AddIntParam("NewIndex", index).
		Do()
	if err != nil {
		return tr064model.GetGenericDeviceInfosResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetGenericDeviceInfosResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetGenericDeviceInfosResponse{}, err
	}
	return result, nil
}
