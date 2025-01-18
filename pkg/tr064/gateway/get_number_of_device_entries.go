package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetNumberOfDeviceEntries AUTO-GENERATED (do not edit) code from [x_homeplugSCPD],
// based on SOAP action 'GetNumberOfDeviceEntries', Fritz!Box-System-Version 164.08.00
//
// [x_homeplugSCPD]: http://fritz.box:49000/x_homeplugSCPD.xml
func GetNumberOfDeviceEntries(session *soap.SoapSession) (tr064model.GetNumberOfDeviceEntriesResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_homeplug").
		Uri("urn:dslforum-org:service:X_AVM-DE_Homeplug:1").
		Action("GetNumberOfDeviceEntries").
		Do()
	if err != nil {
		return tr064model.GetNumberOfDeviceEntriesResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetNumberOfDeviceEntriesResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetNumberOfDeviceEntriesResponse{}, err
	}
	return result, nil
}
