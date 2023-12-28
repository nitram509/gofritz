package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetDefaultConnectionService AUTO-GENERATED (do not edit) code from [layer3forwardingSCPD],
// based on SOAP action 'SetDefaultConnectionService', Fritz!Box-System-Version 164.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
func SetDefaultConnectionService(session *soap.SoapSession, defaultConnectionService string) (tr064model.SetDefaultConnectionServiceResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/layer3forwarding").
		Uri("urn:dslforum-org:service:Layer3Forwarding:1").
		Action("SetDefaultConnectionService").
		AddStringParam("NewDefaultConnectionService", defaultConnectionService).
		Do().Body.Data
	result := tr064model.SetDefaultConnectionServiceResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
