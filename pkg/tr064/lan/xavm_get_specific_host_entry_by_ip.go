package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetSpecificHostEntryByIP AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetSpecificHostEntryByIP', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmGetSpecificHostEntryByIP(session *soap.SoapSession, ipAddress string) (tr064model.XavmGetSpecificHostEntryByIPResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetSpecificHostEntryByIP").
		AddStringParam("NewIPAddress", ipAddress).
		Do()
	if err != nil {
		return tr064model.XavmGetSpecificHostEntryByIPResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetSpecificHostEntryByIPResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetSpecificHostEntryByIPResponse{}, err
	}
	return result, nil
}
