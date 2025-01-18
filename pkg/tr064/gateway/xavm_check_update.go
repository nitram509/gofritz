package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmCheckUpdate AUTO-GENERATED (do not edit) code from [userifSCPD],
// based on SOAP action 'X_AVM-DE_CheckUpdate', Fritz!Box-System-Version 164.08.00
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
func XavmCheckUpdate(session *soap.SoapSession) (tr064model.XavmCheckUpdateResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/userif").
		Uri("urn:dslforum-org:service:UserInterface:1").
		Action("X_AVM-DE_CheckUpdate").
		Do()
	if err != nil {
		return tr064model.XavmCheckUpdateResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmCheckUpdateResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmCheckUpdateResponse{}, err
	}
	return result, nil
}
