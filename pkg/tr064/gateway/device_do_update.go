package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// DeviceDoUpdate AUTO-GENERATED (do not edit) code from [x_homeplugSCPD],
// based on SOAP action 'DeviceDoUpdate', Fritz!Box-System-Version 164.07.57
//
// [x_homeplugSCPD]: http://fritz.box:49000/x_homeplugSCPD.xml
func DeviceDoUpdate(session *soap.SoapSession, macAddress string) (tr064model.DeviceDoUpdateResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_homeplug").
		Uri("urn:dslforum-org:service:X_AVM-DE_Homeplug:1").
		Action("DeviceDoUpdate").
		AddStringParam("NewMACAddress", macAddress).
		Do()
	if err != nil {
		return tr064model.DeviceDoUpdateResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.DeviceDoUpdateResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.DeviceDoUpdateResponse{}, err
	}
	return result, nil
}
