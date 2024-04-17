package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmDoPrepareCGI AUTO-GENERATED (do not edit) code from [userifSCPD],
// based on SOAP action 'X_AVM-DE_DoPrepareCGI', Fritz!Box-System-Version 164.07.57
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
func XavmDoPrepareCGI(session *soap.SoapSession) (tr064model.XavmDoPrepareCGIResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/userif").
		Uri("urn:dslforum-org:service:UserInterface:1").
		Action("X_AVM-DE_DoPrepareCGI").
		Do().Body.Data
	result := tr064model.XavmDoPrepareCGIResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
