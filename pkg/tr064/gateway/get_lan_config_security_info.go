package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetLanConfigSecurityInfo AUTO-GENERATED (do not edit) code from [lanconfigsecuritySCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 141.07.57
//
// [lanconfigsecuritySCPD]: http://fritz.box:49000/lanconfigsecuritySCPD.xml
func GetLanConfigSecurityInfo(session *soap.SoapSession) (tr064model.GetLanConfigSecurityInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanconfigsecurity").
		Uri("urn:dslforum-org:service:LANConfigSecurity:1").
		Action("GetInfo").
		Do().Body.Data
	result := tr064model.GetLanConfigSecurityInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
