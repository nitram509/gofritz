package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetInternationalConfig AUTO-GENERATED (do not edit) code from [userifSCPD],
// based on SOAP action 'X_AVM-DE_GetInternationalConfig', Fritz!Box-System-Version 164.08.00
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
func XavmGetInternationalConfig(session *soap.SoapSession) (tr064model.XavmGetInternationalConfigResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/userif").
		Uri("urn:dslforum-org:service:UserInterface:1").
		Action("X_AVM-DE_GetInternationalConfig").
		Do()
	if err != nil {
		return tr064model.XavmGetInternationalConfigResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetInternationalConfigResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetInternationalConfigResponse{}, err
	}
	return result, nil
}
