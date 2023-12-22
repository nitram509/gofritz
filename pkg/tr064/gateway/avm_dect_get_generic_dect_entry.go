package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetGenericDectEntry AUTO-GENERATED (do not edit) code from [x_dectSCPD],
// based on SOAP action 'GetGenericDectEntry', Fritz!Box-System-Version 164.07.57
//
// [x_dectSCPD]: http://fritz.box:49000/x_dectSCPD.xml
func GetGenericDectEntry(session *soap.SoapSession, index int) (tr064model.GetGenericDectEntryResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_dect").
		Uri("urn:dslforum-org:service:X_AVM-DE_Dect:1").
		Action("GetGenericDectEntry").
		AddIntParam("NewIndex", index).
		Do().Body.Data
	result := tr064model.GetGenericDectEntryResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
