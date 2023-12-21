package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetUserInterfaceInfo AUTO-GENERATED (do not edit) code from [userifSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
func GetUserInterfaceInfo(session *soap.SoapSession) (tr064model.GetUserInterfaceInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/userif").
		Uri("urn:dslforum-org:service:UserInterface:1").
		Action("GetInfo").
		Do().Body.Data
	result := tr064model.GetUserInterfaceInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
