package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetDNSServers AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetDNSServers', Fritz!Box-System-Version 164.08.00
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func GetDNSServers(session *soap.SoapSession) (tr064model.GetDNSServersResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("GetDNSServers").
		Do()
	if err != nil {
		return tr064model.GetDNSServersResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetDNSServersResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetDNSServersResponse{}, err
	}
	return result, nil
}
