package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetWlanConfigurationConfig AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func SetWlanConfigurationConfig(session *soap.SoapSession) (tr064model.SetWlanConfigurationConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("SetConfig").
		Do().Body.Data
	result := tr064model.SetWlanConfigurationConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
