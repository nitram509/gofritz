package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmHostDoUpdate AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_HostDoUpdate', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmHostDoUpdate(session *soap.SoapSession, macAddress string) (tr064model.XavmHostDoUpdateResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_HostDoUpdate").
		AddStringParam("NewMACAddress", macAddress).
		Do()
	if err != nil {
		return tr064model.XavmHostDoUpdateResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmHostDoUpdateResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmHostDoUpdateResponse{}, err
	}
	return result, nil
}
