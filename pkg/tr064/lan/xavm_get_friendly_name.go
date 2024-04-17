package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetFriendlyName AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetFriendlyName', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmGetFriendlyName(session *soap.SoapSession) (tr064model.XavmGetFriendlyNameResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetFriendlyName").
		Do()
	if err != nil {
		return tr064model.XavmGetFriendlyNameResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetFriendlyNameResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetFriendlyNameResponse{}, err
	}
	return result, nil
}
