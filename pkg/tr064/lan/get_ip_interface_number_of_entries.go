package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetIPInterfaceNumberOfEntries AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetIPInterfaceNumberOfEntries', Fritz!Box-System-Version 164.08.00
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func GetIPInterfaceNumberOfEntries(session *soap.SoapSession) (tr064model.GetIPInterfaceNumberOfEntriesResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("GetIPInterfaceNumberOfEntries").
		Do()
	if err != nil {
		return tr064model.GetIPInterfaceNumberOfEntriesResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetIPInterfaceNumberOfEntriesResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetIPInterfaceNumberOfEntriesResponse{}, err
	}
	return result, nil
}
