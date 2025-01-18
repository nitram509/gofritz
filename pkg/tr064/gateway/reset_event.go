package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// ResetEvent AUTO-GENERATED (do not edit) code from [x_appsetupSCPD],
// based on SOAP action 'ResetEvent', Fritz!Box-System-Version 164.08.00
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
func ResetEvent(session *soap.SoapSession, eventId int) (tr064model.ResetEventResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_appsetup").
		Uri("urn:dslforum-org:service:X_AVM-DE_AppSetup:1").
		Action("ResetEvent").
		AddIntParam("NewEventId", eventId).
		Do()
	if err != nil {
		return tr064model.ResetEventResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.ResetEventResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.ResetEventResponse{}, err
	}
	return result, nil
}
