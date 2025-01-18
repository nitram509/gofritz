package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetNTPServers AUTO-GENERATED (do not edit) code from [timeSCPD],
// based on SOAP action 'SetNTPServers', Fritz!Box-System-Version 164.08.00
//
// [timeSCPD]: http://fritz.box:49000/timeSCPD.xml
func SetNTPServers(session *soap.SoapSession, ntpServer1 string, ntpServer2 string) (tr064model.SetNTPServersResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/time").
		Uri("urn:dslforum-org:service:Time:1").
		Action("SetNTPServers").
		AddStringParam("NewNTPServer1", ntpServer1).
		AddStringParam("NewNTPServer2", ntpServer2).
		Do()
	if err != nil {
		return tr064model.SetNTPServersResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetNTPServersResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetNTPServersResponse{}, err
	}
	return result, nil
}
