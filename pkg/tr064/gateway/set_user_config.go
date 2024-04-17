package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetUserConfig AUTO-GENERATED (do not edit) code from [x_storageSCPD],
// based on SOAP action 'SetUserConfig', Fritz!Box-System-Version 164.07.57
//
// [x_storageSCPD]: http://fritz.box:49000/x_storageSCPD.xml
func SetUserConfig(session *soap.SoapSession, enable bool, password string, avmNetworkAccessReadOnly bool) (tr064model.SetUserConfigResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_storage").
		Uri("urn:dslforum-org:service:X_AVM-DE_Storage:1").
		Action("SetUserConfig").
		AddBoolParam("NewEnable", enable).
		AddStringParam("NewPassword", password).
		AddBoolParam("NewX_AVM-DE_NetworkAccessReadOnly", avmNetworkAccessReadOnly).
		Do()
	if err != nil {
		return tr064model.SetUserConfigResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetUserConfigResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetUserConfigResponse{}, err
	}
	return result, nil
}
