package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetChangeCounter AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetChangeCounter', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmGetChangeCounter(session *soap.SoapSession) (tr064model.XavmGetChangeCounterResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetChangeCounter").
		Do()
	if err != nil {
		return tr064model.XavmGetChangeCounterResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetChangeCounterResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetChangeCounterResponse{}, err
	}
	return result, nil
}
