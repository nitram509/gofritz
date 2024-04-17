package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// ResetEvent AUTO-GENERATED (do not edit) code from [x_appsetupSCPD],
// based on SOAP action 'ResetEvent', Fritz!Box-System-Version 164.07.57
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
func ResetEvent(session *soap.SoapSession, eventId int) (tr064model.ResetEventResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_appsetup").
		Uri("urn:dslforum-org:service:X_AVM-DE_AppSetup:1").
		Action("ResetEvent").
		AddIntParam("NewEventId", eventId).
		Do().Body.Data
	result := tr064model.ResetEventResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
