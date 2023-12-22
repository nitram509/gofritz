package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetSpecificDectEntry AUTO-GENERATED (do not edit) code from [x_dectSCPD],
// based on SOAP action 'GetSpecificDectEntry', Fritz!Box-System-Version 164.07.57
//
// [x_dectSCPD]: http://fritz.box:49000/x_dectSCPD.xml
func GetSpecificDectEntry(session *soap.SoapSession, id string) (tr064model.GetSpecificDectEntryResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_dect").
		Uri("urn:dslforum-org:service:X_AVM-DE_Dect:1").
		Action("GetSpecificDectEntry").
		AddStringParam("NewID", id).
		Do().Body.Data
	result := tr064model.GetSpecificDectEntryResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
