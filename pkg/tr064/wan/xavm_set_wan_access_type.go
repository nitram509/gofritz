package wan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetWANAccessType AUTO-GENERATED (do not edit) code from [wancommonifconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetWANAccessType', Fritz!Box-System-Version 164.08.00
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
func XavmSetWANAccessType(session *soap.SoapSession, accessType string) (tr064model.XavmSetWANAccessTypeResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wancommonifconfig1").
		Uri("urn:dslforum-org:service:WANCommonInterfaceConfig:1").
		Action("X_AVM-DE_SetWANAccessType").
		AddStringParam("NewAccessType", accessType).
		Do()
	if err != nil {
		return tr064model.XavmSetWANAccessTypeResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetWANAccessTypeResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetWANAccessTypeResponse{}, err
	}
	return result, nil
}
