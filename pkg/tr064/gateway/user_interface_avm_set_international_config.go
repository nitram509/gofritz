package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmSetInternationalConfig AUTO-GENERATED (do not edit) code from [userifSCPD],
// based on SOAP action 'X_AVM-DE_SetInternationalConfig', Fritz!Box-System-Version 164.07.57
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
func XavmSetInternationalConfig(session *soap.SoapSession) (tr064model.XavmSetInternationalConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/userif").
		Uri("urn:dslforum-org:service:UserInterface:1").
		Action("X_AVM-DE_SetInternationalConfig").
		Do().Body.Data
	result := tr064model.XavmSetInternationalConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
