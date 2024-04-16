package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetUSPMyFRITZEnable AUTO-GENERATED (do not edit) code from [x_uspcontrollerSCPD],
// based on SOAP action 'GetUSPMyFRITZEnable', Fritz!Box-System-Version 164.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
func GetUSPMyFRITZEnable(session *soap.SoapSession) (tr064model.GetUSPMyFRITZEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_uspcontroller").
		Uri("urn:dslforum-org:service:X_AVM-DE_USPController:1").
		Action("GetUSPMyFRITZEnable").
		Do().Body.Data
	result := tr064model.GetUSPMyFRITZEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
