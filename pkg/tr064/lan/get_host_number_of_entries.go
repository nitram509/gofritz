package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetHostNumberOfEntries AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'GetHostNumberOfEntries', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func GetHostNumberOfEntries(session *soap.SoapSession) (tr064model.GetHostNumberOfEntriesResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("GetHostNumberOfEntries").
		Do()
	if err != nil {
		return tr064model.GetHostNumberOfEntriesResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetHostNumberOfEntriesResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetHostNumberOfEntriesResponse{}, err
	}
	return result, nil
}
