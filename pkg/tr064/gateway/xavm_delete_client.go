package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmDeleteClient AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_DeleteClient', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmDeleteClient(session *soap.SoapSession, avmClientIndex int) (tr064model.XavmDeleteClientResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_DeleteClient").
		AddIntParam("NewX_AVM-DE_ClientIndex", avmClientIndex).
		Do()
	if err != nil {
		return tr064model.XavmDeleteClientResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmDeleteClientResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmDeleteClientResponse{}, err
	}
	return result, nil
}
