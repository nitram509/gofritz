package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetState AUTO-GENERATED (do not edit) code from [x_authSCPD],
// based on SOAP action 'GetState', Fritz!Box-System-Version 164.07.57
//
// [x_authSCPD]: http://fritz.box:49000/x_authSCPD.xml
func GetState(session *soap.SoapSession) (tr064model.GetStateResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_auth").
		Uri("urn:dslforum-org:service:X_AVM-DE_Auth:1").
		Action("GetState").
		Do().Body.Data
	result := tr064model.GetStateResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
