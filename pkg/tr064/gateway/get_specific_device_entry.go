package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetSpecificDeviceEntry AUTO-GENERATED (do not edit) code from [x_homeplugSCPD],
// based on SOAP action 'GetSpecificDeviceEntry', Fritz!Box-System-Version 164.07.57
//
// [x_homeplugSCPD]: http://fritz.box:49000/x_homeplugSCPD.xml
func GetSpecificDeviceEntry(session *soap.SoapSession, macAddress string) (tr064model.GetSpecificDeviceEntryResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_homeplug").
		Uri("urn:dslforum-org:service:X_AVM-DE_Homeplug:1").
		Action("GetSpecificDeviceEntry").
		AddStringParam("NewMACAddress", macAddress).
		Do().Body.Data
	result := tr064model.GetSpecificDeviceEntryResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
