package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetAvmAuthConfig AUTO-GENERATED (do not edit) code from [x_authSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.07.57
//
// [x_authSCPD]: http://fritz.box:49000/x_authSCPD.xml
func SetAvmAuthConfig(session *soap.SoapSession, action string) (tr064model.SetAvmAuthConfigResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_auth").
		Uri("urn:dslforum-org:service:X_AVM-DE_Auth:1").
		Action("SetConfig").
		AddStringParam("NewAction", action).
		Do()
	if err != nil {
		return tr064model.SetAvmAuthConfigResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetAvmAuthConfigResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetAvmAuthConfigResponse{}, err
	}
	return result, nil
}
