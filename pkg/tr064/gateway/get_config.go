package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetConfig AUTO-GENERATED (do not edit) code from [x_appsetupSCPD],
// based on SOAP action 'GetConfig', Fritz!Box-System-Version 164.07.57
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
func GetConfig(session *soap.SoapSession) (tr064model.GetConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_appsetup").
		Uri("urn:dslforum-org:service:X_AVM-DE_AppSetup:1").
		Action("GetConfig").
		Do().Body.Data
	result := tr064model.GetConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
