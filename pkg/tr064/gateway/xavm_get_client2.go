package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetClient2 AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetClient2', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmGetClient2(session *soap.SoapSession, avmClientIndex int) (tr064model.XavmGetClient2Response, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_GetClient2").
		AddIntParam("NewX_AVM-DE_ClientIndex", avmClientIndex).
		Do().Body.Data
	result := tr064model.XavmGetClient2Response{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
