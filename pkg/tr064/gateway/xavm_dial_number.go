package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmDialNumber AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_DialNumber', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmDialNumber(session *soap.SoapSession, avmPhoneNumber string) (tr064model.XavmDialNumberResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_DialNumber").
		AddStringParam("NewX_AVM-DE_PhoneNumber", avmPhoneNumber).
		Do()
	if err != nil {
		return tr064model.XavmDialNumberResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmDialNumberResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmDialNumberResponse{}, err
	}
	return result, nil
}
