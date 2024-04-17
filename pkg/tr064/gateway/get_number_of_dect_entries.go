package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetNumberOfDectEntries AUTO-GENERATED (do not edit) code from [x_dectSCPD],
// based on SOAP action 'GetNumberOfDectEntries', Fritz!Box-System-Version 164.07.57
//
// [x_dectSCPD]: http://fritz.box:49000/x_dectSCPD.xml
func GetNumberOfDectEntries(session *soap.SoapSession) (tr064model.GetNumberOfDectEntriesResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_dect").
		Uri("urn:dslforum-org:service:X_AVM-DE_Dect:1").
		Action("GetNumberOfDectEntries").
		Do().Body.Data
	result := tr064model.GetNumberOfDectEntriesResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
