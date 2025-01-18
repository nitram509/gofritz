package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetUserInterfaceInfo AUTO-GENERATED (do not edit) code from [userifSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
func GetUserInterfaceInfo(session *soap.SoapSession) (tr064model.GetUserInterfaceInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/userif").
		Uri("urn:dslforum-org:service:UserInterface:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetUserInterfaceInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetUserInterfaceInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetUserInterfaceInfoResponse{}, err
	}
	return result, nil
}
