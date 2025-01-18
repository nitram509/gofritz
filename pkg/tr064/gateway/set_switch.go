package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetSwitch AUTO-GENERATED (do not edit) code from [x_homeautoSCPD],
// based on SOAP action 'SetSwitch', Fritz!Box-System-Version 164.08.00
//
// [x_homeautoSCPD]: http://fritz.box:49000/x_homeautoSCPD.xml
func SetSwitch(session *soap.SoapSession, aIN string, switchState string) (tr064model.SetSwitchResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_homeauto").
		Uri("urn:dslforum-org:service:X_AVM-DE_Homeauto:1").
		Action("SetSwitch").
		AddStringParam("NewAIN", aIN).
		AddStringParam("NewSwitchState", switchState).
		Do()
	if err != nil {
		return tr064model.SetSwitchResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetSwitchResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetSwitchResponse{}, err
	}
	return result, nil
}
