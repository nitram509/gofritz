package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmSetPrioritizationByIP AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_SetPrioritizationByIP', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmSetPrioritizationByIP(session *soap.SoapSession, ipAddress string, avmPriority bool) (tr064model.XavmSetPrioritizationByIPResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_SetPrioritizationByIP").
		AddStringParam("NewIPAddress", ipAddress).
		AddBoolParam("NewX_AVM-DE_Priority", avmPriority).
		Do().Body.Data
	result := tr064model.XavmSetPrioritizationByIPResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
