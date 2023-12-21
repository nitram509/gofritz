package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetSwitch AUTO-GENERATED (do not edit) code from [x_homeautoSCPD],
// based on SOAP action 'SetSwitch', Fritz!Box-System-Version 164.07.57
//
// [x_homeautoSCPD]: http://fritz.box:49000/x_homeautoSCPD.xml
func SetSwitch(session *soap.SoapSession) (tr064model.SetSwitchResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_homeauto").
		Uri("urn:dslforum-org:service:X_AVM-DE_Homeauto:1").
		Action("SetSwitch").
		Do().Body.Data
	result := tr064model.SetSwitchResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
