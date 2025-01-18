package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetGenericForwardingEntry AUTO-GENERATED (do not edit) code from [layer3forwardingSCPD],
// based on SOAP action 'GetGenericForwardingEntry', Fritz!Box-System-Version 164.08.00
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
func GetGenericForwardingEntry(session *soap.SoapSession, forwardingIndex int) (tr064model.GetGenericForwardingEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/layer3forwarding").
		Uri("urn:dslforum-org:service:Layer3Forwarding:1").
		Action("GetGenericForwardingEntry").
		AddIntParam("NewForwardingIndex", forwardingIndex).
		Do()
	if err != nil {
		return tr064model.GetGenericForwardingEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetGenericForwardingEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetGenericForwardingEntryResponse{}, err
	}
	return result, nil
}
