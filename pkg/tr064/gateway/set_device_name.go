package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetDeviceName AUTO-GENERATED (do not edit) code from [x_homeautoSCPD],
// based on SOAP action 'SetDeviceName', Fritz!Box-System-Version 164.07.57
//
// [x_homeautoSCPD]: http://fritz.box:49000/x_homeautoSCPD.xml
func SetDeviceName(session *soap.SoapSession, aIN string, deviceName string) (tr064model.SetDeviceNameResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_homeauto").
		Uri("urn:dslforum-org:service:X_AVM-DE_Homeauto:1").
		Action("SetDeviceName").
		AddStringParam("NewAIN", aIN).
		AddStringParam("NewDeviceName", deviceName).
		Do()
	if err != nil {
		return tr064model.SetDeviceNameResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetDeviceNameResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetDeviceNameResponse{}, err
	}
	return result, nil
}
