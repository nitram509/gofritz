package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetLanConfigSecurityInfo AUTO-GENERATED (do not edit) code from [lanconfigsecuritySCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [lanconfigsecuritySCPD]: http://fritz.box:49000/lanconfigsecuritySCPD.xml
func GetLanConfigSecurityInfo(session *soap.SoapSession) (tr064model.GetLanConfigSecurityInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanconfigsecurity").
		Uri("urn:dslforum-org:service:LANConfigSecurity:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetLanConfigSecurityInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetLanConfigSecurityInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetLanConfigSecurityInfoResponse{}, err
	}
	return result, nil
}
