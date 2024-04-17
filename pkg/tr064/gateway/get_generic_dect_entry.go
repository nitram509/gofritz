package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetGenericDectEntry AUTO-GENERATED (do not edit) code from [x_dectSCPD],
// based on SOAP action 'GetGenericDectEntry', Fritz!Box-System-Version 164.07.57
//
// [x_dectSCPD]: http://fritz.box:49000/x_dectSCPD.xml
func GetGenericDectEntry(session *soap.SoapSession, index int) (tr064model.GetGenericDectEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_dect").
		Uri("urn:dslforum-org:service:X_AVM-DE_Dect:1").
		Action("GetGenericDectEntry").
		AddIntParam("NewIndex", index).
		Do()
	if err != nil {
		return tr064model.GetGenericDectEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetGenericDectEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetGenericDectEntryResponse{}, err
	}
	return result, nil
}
