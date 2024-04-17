package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetMyFRITZ AUTO-GENERATED (do not edit) code from [x_myfritzSCPD],
// based on SOAP action 'SetMyFRITZ', Fritz!Box-System-Version 164.07.57
//
// [x_myfritzSCPD]: http://fritz.box:49000/x_myfritzSCPD.xml
func SetMyFRITZ(session *soap.SoapSession, enabled bool, email string) (tr064model.SetMyFRITZResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_myfritz").
		Uri("urn:dslforum-org:service:X_AVM-DE_MyFritz:1").
		Action("SetMyFRITZ").
		AddBoolParam("NewEnabled", enabled).
		AddStringParam("NewEmail", email).
		Do()
	if err != nil {
		return tr064model.SetMyFRITZResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetMyFRITZResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetMyFRITZResponse{}, err
	}
	return result, nil
}
