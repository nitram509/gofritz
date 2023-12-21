package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetServiceByIndex AUTO-GENERATED (do not edit) code from [x_myfritzSCPD],
// based on SOAP action 'SetServiceByIndex', Fritz!Box-System-Version 164.07.57
//
// [x_myfritzSCPD]: http://fritz.box:49000/x_myfritzSCPD.xml
func SetServiceByIndex(session *soap.SoapSession) (tr064model.SetServiceByIndexResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_myfritz").
		Uri("urn:dslforum-org:service:X_AVM-DE_MyFritz:1").
		Action("SetServiceByIndex").
		Do().Body.Data
	result := tr064model.SetServiceByIndexResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
