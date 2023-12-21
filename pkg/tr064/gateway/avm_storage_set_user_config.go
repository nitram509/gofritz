package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetUserConfig AUTO-GENERATED (do not edit) code from [x_storageSCPD],
// based on SOAP action 'SetUserConfig', Fritz!Box-System-Version 164.07.57
//
// [x_storageSCPD]: http://fritz.box:49000/x_storageSCPD.xml
func SetUserConfig(session *soap.SoapSession) (tr064model.SetUserConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_storage").
		Uri("urn:dslforum-org:service:X_AVM-DE_Storage:1").
		Action("SetUserConfig").
		Do().Body.Data
	result := tr064model.SetUserConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
