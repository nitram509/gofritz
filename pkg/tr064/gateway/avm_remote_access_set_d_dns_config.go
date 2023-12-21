package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetDDNSConfig AUTO-GENERATED (do not edit) code from [x_remoteSCPD],
// based on SOAP action 'SetDDNSConfig', Fritz!Box-System-Version 164.07.57
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
func SetDDNSConfig(session *soap.SoapSession) (tr064model.SetDDNSConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_remote").
		Uri("urn:dslforum-org:service:X_AVM-DE_RemoteAccess:1").
		Action("SetDDNSConfig").
		Do().Body.Data
	result := tr064model.SetDDNSConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
