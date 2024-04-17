package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAvmAuthInfo AUTO-GENERATED (do not edit) code from [x_authSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_authSCPD]: http://fritz.box:49000/x_authSCPD.xml
func GetAvmAuthInfo(session *soap.SoapSession) (tr064model.GetAvmAuthInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_auth").
		Uri("urn:dslforum-org:service:X_AVM-DE_Auth:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetAvmAuthInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetAvmAuthInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetAvmAuthInfoResponse{}, err
	}
	return result, nil
}
