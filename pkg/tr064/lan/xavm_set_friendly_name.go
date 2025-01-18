package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetFriendlyName AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_SetFriendlyName', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmSetFriendlyName(session *soap.SoapSession, avmFriendlyName string) (tr064model.XavmSetFriendlyNameResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_SetFriendlyName").
		AddStringParam("NewX_AVM-DE_FriendlyName", avmFriendlyName).
		Do()
	if err != nil {
		return tr064model.XavmSetFriendlyNameResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetFriendlyNameResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetFriendlyNameResponse{}, err
	}
	return result, nil
}
