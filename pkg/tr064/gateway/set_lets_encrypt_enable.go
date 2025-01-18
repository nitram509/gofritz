package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetLetsEncryptEnable AUTO-GENERATED (do not edit) code from [x_remoteSCPD],
// based on SOAP action 'SetLetsEncryptEnable', Fritz!Box-System-Version 164.08.00
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
func SetLetsEncryptEnable(session *soap.SoapSession, letsEncryptEnabled bool) (tr064model.SetLetsEncryptEnableResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_remote").
		Uri("urn:dslforum-org:service:X_AVM-DE_RemoteAccess:1").
		Action("SetLetsEncryptEnable").
		AddBoolParam("NewLetsEncryptEnabled", letsEncryptEnabled).
		Do()
	if err != nil {
		return tr064model.SetLetsEncryptEnableResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetLetsEncryptEnableResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetLetsEncryptEnableResponse{}, err
	}
	return result, nil
}
