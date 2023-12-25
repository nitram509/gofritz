package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmGetAlarmClock AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetAlarmClock', Fritz!Box-System-Version 141.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmGetAlarmClock(session *soap.SoapSession, index int) (tr064model.XavmGetAlarmClockResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_GetAlarmClock").
		AddIntParam("NewIndex", index).
		Do().Body.Data
	result := tr064model.XavmGetAlarmClockResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
