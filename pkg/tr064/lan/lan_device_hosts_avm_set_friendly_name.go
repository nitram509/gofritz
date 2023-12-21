package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmSetFriendlyName AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_SetFriendlyName', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmSetFriendlyName(session *soap.SoapSession) (tr064model.XavmSetFriendlyNameResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_SetFriendlyName").
		Do().Body.Data
	result := tr064model.XavmSetFriendlyNameResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
