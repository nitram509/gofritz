package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetAvmRemoteAccessEnable AUTO-GENERATED (do not edit) code from [x_remoteSCPD],
// based on SOAP action 'SetEnable', Fritz!Box-System-Version 164.07.57
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
func SetAvmRemoteAccessEnable(session *soap.SoapSession, enabled bool) (tr064model.SetAvmRemoteAccessEnableResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_remote").
		Uri("urn:dslforum-org:service:X_AVM-DE_RemoteAccess:1").
		Action("SetEnable").
		AddBoolParam("NewEnabled", enabled).
		Do()
	if err != nil {
		return tr064model.SetAvmRemoteAccessEnableResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetAvmRemoteAccessEnableResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetAvmRemoteAccessEnableResponse{}, err
	}
	return result, nil
}
