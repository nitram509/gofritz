package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetPrioritizationByIP AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_SetPrioritizationByIP', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmSetPrioritizationByIP(session *soap.SoapSession, ipAddress string, avmPriority bool) (tr064model.XavmSetPrioritizationByIPResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_SetPrioritizationByIP").
		AddStringParam("NewIPAddress", ipAddress).
		AddBoolParam("NewX_AVM-DE_Priority", avmPriority).
		Do()
	if err != nil {
		return tr064model.XavmSetPrioritizationByIPResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetPrioritizationByIPResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetPrioritizationByIPResponse{}, err
	}
	return result, nil
}
