package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetInternationalConfig AUTO-GENERATED (do not edit) code from [userifSCPD],
// based on SOAP action 'X_AVM-DE_SetInternationalConfig', Fritz!Box-System-Version 164.08.00
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
func XavmSetInternationalConfig(session *soap.SoapSession, avmLanguage string, avmCountry string, avmAnnex string) (tr064model.XavmSetInternationalConfigResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/userif").
		Uri("urn:dslforum-org:service:UserInterface:1").
		Action("X_AVM-DE_SetInternationalConfig").
		AddStringParam("NewX_AVM-DE_Language", avmLanguage).
		AddStringParam("NewX_AVM-DE_Country", avmCountry).
		AddStringParam("NewX_AVM-DE_Annex", avmAnnex).
		Do()
	if err != nil {
		return tr064model.XavmSetInternationalConfigResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetInternationalConfigResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetInternationalConfigResponse{}, err
	}
	return result, nil
}
