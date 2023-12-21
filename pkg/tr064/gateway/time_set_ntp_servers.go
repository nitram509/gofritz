package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetNTPServers AUTO-GENERATED (do not edit) code from [timeSCPD],
// based on SOAP action 'SetNTPServers', Fritz!Box-System-Version 164.07.57
//
// [timeSCPD]: http://fritz.box:49000/timeSCPD.xml
func SetNTPServers(session *soap.SoapSession) (tr064model.SetNTPServersResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/time").
		Uri("urn:dslforum-org:service:Time:1").
		Action("SetNTPServers").
		Do().Body.Data
	result := tr064model.SetNTPServersResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
