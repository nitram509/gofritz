package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetAvmRemoteAccessConfig AUTO-GENERATED (do not edit) code from [x_remoteSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.07.57
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
func SetAvmRemoteAccessConfig(session *soap.SoapSession, enabled bool, port string, username string, password string) (tr064model.SetAvmRemoteAccessConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_remote").
		Uri("urn:dslforum-org:service:X_AVM-DE_RemoteAccess:1").
		Action("SetConfig").
		AddBoolParam("NewEnabled", enabled).
		AddStringParam("NewPort", port).
		AddStringParam("NewUsername", username).
		AddStringParam("NewPassword", password).
		Do().Body.Data
	result := tr064model.SetAvmRemoteAccessConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
