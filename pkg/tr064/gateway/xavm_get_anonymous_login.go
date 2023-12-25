package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmGetAnonymousLogin AUTO-GENERATED (do not edit) code from [lanconfigsecuritySCPD],
// based on SOAP action 'X_AVM-DE_GetAnonymousLogin', Fritz!Box-System-Version 141.07.57
//
// [lanconfigsecuritySCPD]: http://fritz.box:49000/lanconfigsecuritySCPD.xml
func XavmGetAnonymousLogin(session *soap.SoapSession) (tr064model.XavmGetAnonymousLoginResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanconfigsecurity").
		Uri("urn:dslforum-org:service:LANConfigSecurity:1").
		Action("X_AVM-DE_GetAnonymousLogin").
		Do().Body.Data
	result := tr064model.XavmGetAnonymousLoginResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
