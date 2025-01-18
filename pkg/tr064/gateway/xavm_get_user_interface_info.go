package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetUserInterfaceInfo AUTO-GENERATED (do not edit) code from [userifSCPD],
// based on SOAP action 'X_AVM-DE_GetInfo', Fritz!Box-System-Version 164.08.00
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
func XavmGetUserInterfaceInfo(session *soap.SoapSession) (tr064model.XavmGetUserInterfaceInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/userif").
		Uri("urn:dslforum-org:service:UserInterface:1").
		Action("X_AVM-DE_GetInfo").
		Do()
	if err != nil {
		return tr064model.XavmGetUserInterfaceInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetUserInterfaceInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetUserInterfaceInfoResponse{}, err
	}
	return result, nil
}
