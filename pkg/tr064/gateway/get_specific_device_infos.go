package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetSpecificDeviceInfos AUTO-GENERATED (do not edit) code from [x_homeautoSCPD],
// based on SOAP action 'GetSpecificDeviceInfos', Fritz!Box-System-Version 164.08.00
//
// [x_homeautoSCPD]: http://fritz.box:49000/x_homeautoSCPD.xml
func GetSpecificDeviceInfos(session *soap.SoapSession, aIN string) (tr064model.GetSpecificDeviceInfosResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_homeauto").
		Uri("urn:dslforum-org:service:X_AVM-DE_Homeauto:1").
		Action("GetSpecificDeviceInfos").
		AddStringParam("NewAIN", aIN).
		Do()
	if err != nil {
		return tr064model.GetSpecificDeviceInfosResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetSpecificDeviceInfosResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetSpecificDeviceInfosResponse{}, err
	}
	return result, nil
}
