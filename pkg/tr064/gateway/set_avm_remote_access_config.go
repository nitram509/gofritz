package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetAvmRemoteAccessConfig AUTO-GENERATED (do not edit) code from [x_remoteSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.08.00
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
func SetAvmRemoteAccessConfig(session *soap.SoapSession, enabled bool, port string, username string, password string) (tr064model.SetAvmRemoteAccessConfigResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_remote").
		Uri("urn:dslforum-org:service:X_AVM-DE_RemoteAccess:1").
		Action("SetConfig").
		AddBoolParam("NewEnabled", enabled).
		AddStringParam("NewPort", port).
		AddStringParam("NewUsername", username).
		AddStringParam("NewPassword", password).
		Do()
	if err != nil {
		return tr064model.SetAvmRemoteAccessConfigResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetAvmRemoteAccessConfigResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetAvmRemoteAccessConfigResponse{}, err
	}
	return result, nil
}
