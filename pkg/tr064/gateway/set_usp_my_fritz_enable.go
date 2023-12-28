package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetUSPMyFRITZEnable AUTO-GENERATED (do not edit) code from [x_uspcontrollerSCPD],
// based on SOAP action 'SetUSPMyFRITZEnable', Fritz!Box-System-Version 164.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
func SetUSPMyFRITZEnable(session *soap.SoapSession, uspMyFritzEnabled bool) (tr064model.SetUSPMyFRITZEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_uspcontroller").
		Uri("urn:dslforum-org:service:X_AVM-DE_USPController:1").
		Action("SetUSPMyFRITZEnable").
		AddBoolParam("NewUSPMyFRITZEnabled", uspMyFritzEnabled).
		Do().Body.Data
	result := tr064model.SetUSPMyFRITZEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
