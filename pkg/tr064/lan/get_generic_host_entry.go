package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetGenericHostEntry AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'GetGenericHostEntry', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func GetGenericHostEntry(session *soap.SoapSession, index int) (tr064model.GetGenericHostEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("GetGenericHostEntry").
		AddIntParam("NewIndex", index).
		Do()
	if err != nil {
		return tr064model.GetGenericHostEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetGenericHostEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetGenericHostEntryResponse{}, err
	}
	return result, nil
}
