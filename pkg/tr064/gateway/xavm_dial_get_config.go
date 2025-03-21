package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmDialGetConfig AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_DialGetConfig', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmDialGetConfig(session *soap.SoapSession) (tr064model.XavmDialGetConfigResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_DialGetConfig").
		Do()
	if err != nil {
		return tr064model.XavmDialGetConfigResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmDialGetConfigResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmDialGetConfigResponse{}, err
	}
	return result, nil
}
