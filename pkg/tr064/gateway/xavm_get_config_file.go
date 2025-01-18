package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetConfigFile AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetConfigFile', Fritz!Box-System-Version 164.08.00
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func XavmGetConfigFile(session *soap.SoapSession, avmPassword string) (tr064model.XavmGetConfigFileResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("X_AVM-DE_GetConfigFile").
		AddStringParam("NewX_AVM-DE_Password", avmPassword).
		Do()
	if err != nil {
		return tr064model.XavmGetConfigFileResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetConfigFileResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetConfigFileResponse{}, err
	}
	return result, nil
}
