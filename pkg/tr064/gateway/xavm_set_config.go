package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetConfig AUTO-GENERATED (do not edit) code from [userifSCPD],
// based on SOAP action 'X_AVM-DE_SetConfig', Fritz!Box-System-Version 164.07.57
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
func XavmSetConfig(session *soap.SoapSession, avmAutoUpdateMode string) (tr064model.XavmSetConfigResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/userif").
		Uri("urn:dslforum-org:service:UserInterface:1").
		Action("X_AVM-DE_SetConfig").
		AddStringParam("NewX_AVM-DE_AutoUpdateMode", avmAutoUpdateMode).
		Do()
	if err != nil {
		return tr064model.XavmSetConfigResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetConfigResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetConfigResponse{}, err
	}
	return result, nil
}
