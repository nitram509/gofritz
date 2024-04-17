package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetSpecificHostEntry AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'GetSpecificHostEntry', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func GetSpecificHostEntry(session *soap.SoapSession, macAddress string) (tr064model.GetSpecificHostEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("GetSpecificHostEntry").
		AddStringParam("NewMACAddress", macAddress).
		Do()
	if err != nil {
		return tr064model.GetSpecificHostEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetSpecificHostEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetSpecificHostEntryResponse{}, err
	}
	return result, nil
}
