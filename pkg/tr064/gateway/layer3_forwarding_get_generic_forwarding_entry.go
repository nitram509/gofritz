package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetGenericForwardingEntry AUTO-GENERATED (do not edit) code from [layer3forwardingSCPD],
// based on SOAP action 'GetGenericForwardingEntry', Fritz!Box-System-Version 164.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
func GetGenericForwardingEntry(session *soap.SoapSession) (tr064model.GetGenericForwardingEntryResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/layer3forwarding").
		Uri("urn:dslforum-org:service:Layer3Forwarding:1").
		Action("GetGenericForwardingEntry").
		Do().Body.Data
	result := tr064model.GetGenericForwardingEntryResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
