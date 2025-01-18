package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetConfigFile AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetConfigFile', Fritz!Box-System-Version 164.08.00
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func XavmSetConfigFile(session *soap.SoapSession, avmPassword string, avmConfigFileUrl string) (tr064model.XavmSetConfigFileResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("X_AVM-DE_SetConfigFile").
		AddStringParam("NewX_AVM-DE_Password", avmPassword).
		AddStringParam("NewX_AVM-DE_ConfigFileUrl", avmConfigFileUrl).
		Do()
	if err != nil {
		return tr064model.XavmSetConfigFileResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetConfigFileResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetConfigFileResponse{}, err
	}
	return result, nil
}
