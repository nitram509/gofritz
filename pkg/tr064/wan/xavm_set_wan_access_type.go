package wan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmSetWANAccessType AUTO-GENERATED (do not edit) code from [wancommonifconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetWANAccessType', Fritz!Box-System-Version 164.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
func XavmSetWANAccessType(session *soap.SoapSession, accessType string) (tr064model.XavmSetWANAccessTypeResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wancommonifconfig1").
		Uri("urn:dslforum-org:service:WANCommonInterfaceConfig:1").
		Action("X_AVM-DE_SetWANAccessType").
		AddStringParam("NewAccessType", accessType).
		Do().Body.Data
	result := tr064model.XavmSetWANAccessTypeResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
