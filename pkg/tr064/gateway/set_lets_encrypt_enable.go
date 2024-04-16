package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetLetsEncryptEnable AUTO-GENERATED (do not edit) code from [x_remoteSCPD],
// based on SOAP action 'SetLetsEncryptEnable', Fritz!Box-System-Version 164.07.57
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
func SetLetsEncryptEnable(session *soap.SoapSession, letsEncryptEnabled bool) (tr064model.SetLetsEncryptEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_remote").
		Uri("urn:dslforum-org:service:X_AVM-DE_RemoteAccess:1").
		Action("SetLetsEncryptEnable").
		AddBoolParam("NewLetsEncryptEnabled", letsEncryptEnabled).
		Do().Body.Data
	result := tr064model.SetLetsEncryptEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
