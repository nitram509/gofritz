package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetGenericDeviceEntry AUTO-GENERATED (do not edit) code from [x_homeplugSCPD],
// based on SOAP action 'GetGenericDeviceEntry', Fritz!Box-System-Version 164.07.57
//
// [x_homeplugSCPD]: http://fritz.box:49000/x_homeplugSCPD.xml
func GetGenericDeviceEntry(session *soap.SoapSession, index int) (tr064model.GetGenericDeviceEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_homeplug").
		Uri("urn:dslforum-org:service:X_AVM-DE_Homeplug:1").
		Action("GetGenericDeviceEntry").
		AddIntParam("NewIndex", index).
		Do()
	if err != nil {
		return tr064model.GetGenericDeviceEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetGenericDeviceEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetGenericDeviceEntryResponse{}, err
	}
	return result, nil
}
