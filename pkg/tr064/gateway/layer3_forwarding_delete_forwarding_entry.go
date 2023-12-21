package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// DeleteForwardingEntry AUTO-GENERATED (do not edit) code from [layer3forwardingSCPD],
// based on SOAP action 'DeleteForwardingEntry', Fritz!Box-System-Version 164.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
func DeleteForwardingEntry(session *soap.SoapSession) (tr064model.DeleteForwardingEntryResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/layer3forwarding").
		Uri("urn:dslforum-org:service:Layer3Forwarding:1").
		Action("DeleteForwardingEntry").
		Do().Body.Data
	result := tr064model.DeleteForwardingEntryResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
