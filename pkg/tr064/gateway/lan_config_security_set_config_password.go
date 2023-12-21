package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetConfigPassword AUTO-GENERATED (do not edit) code from [lanconfigsecuritySCPD],
// based on SOAP action 'SetConfigPassword', Fritz!Box-System-Version 164.07.57
//
// [lanconfigsecuritySCPD]: http://fritz.box:49000/lanconfigsecuritySCPD.xml
func SetConfigPassword(session *soap.SoapSession) (tr064model.SetConfigPasswordResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanconfigsecurity").
		Uri("urn:dslforum-org:service:LANConfigSecurity:1").
		Action("SetConfigPassword").
		Do().Body.Data
	result := tr064model.SetConfigPasswordResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
