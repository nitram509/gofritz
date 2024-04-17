package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmDoUpdate AUTO-GENERATED (do not edit) code from [userifSCPD],
// based on SOAP action 'X_AVM-DE_DoUpdate', Fritz!Box-System-Version 164.07.57
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
func XavmDoUpdate(session *soap.SoapSession) (tr064model.XavmDoUpdateResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/userif").
		Uri("urn:dslforum-org:service:UserInterface:1").
		Action("X_AVM-DE_DoUpdate").
		Do()
	if err != nil {
		return tr064model.XavmDoUpdateResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmDoUpdateResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmDoUpdateResponse{}, err
	}
	return result, nil
}
