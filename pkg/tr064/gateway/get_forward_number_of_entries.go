package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetForwardNumberOfEntries AUTO-GENERATED (do not edit) code from [layer3forwardingSCPD],
// based on SOAP action 'GetForwardNumberOfEntries', Fritz!Box-System-Version 164.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
func GetForwardNumberOfEntries(session *soap.SoapSession) (tr064model.GetForwardNumberOfEntriesResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/layer3forwarding").
		Uri("urn:dslforum-org:service:Layer3Forwarding:1").
		Action("GetForwardNumberOfEntries").
		Do().Body.Data
	result := tr064model.GetForwardNumberOfEntriesResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
