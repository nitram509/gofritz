package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetNumberOfDeviceEntries AUTO-GENERATED (do not edit) code from [x_homeplugSCPD],
// based on SOAP action 'GetNumberOfDeviceEntries', Fritz!Box-System-Version 141.07.57
//
// [x_homeplugSCPD]: http://fritz.box:49000/x_homeplugSCPD.xml
func GetNumberOfDeviceEntries(session *soap.SoapSession) (tr064model.GetNumberOfDeviceEntriesResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_homeplug").
		Uri("urn:dslforum-org:service:X_AVM-DE_Homeplug:1").
		Action("GetNumberOfDeviceEntries").
		Do().Body.Data
	result := tr064model.GetNumberOfDeviceEntriesResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
