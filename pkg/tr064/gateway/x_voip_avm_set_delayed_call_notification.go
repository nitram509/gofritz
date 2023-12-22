package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmSetDelayedCallNotification AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_SetDelayedCallNotification', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmSetDelayedCallNotification(session *soap.SoapSession, avmClientIndex int, avmDelayedCallNotification bool) (tr064model.XavmSetDelayedCallNotificationResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_SetDelayedCallNotification").
		AddIntParam("NewX_AVM-DE_ClientIndex", avmClientIndex).
		AddBoolParam("NewX_AVM-DE_DelayedCallNotification", avmDelayedCallNotification).
		Do().Body.Data
	result := tr064model.XavmSetDelayedCallNotificationResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
