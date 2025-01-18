package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// DeleteForwardingEntry AUTO-GENERATED (do not edit) code from [layer3forwardingSCPD],
// based on SOAP action 'DeleteForwardingEntry', Fritz!Box-System-Version 164.08.00
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
func DeleteForwardingEntry(session *soap.SoapSession, destIpAddress string, destSubnetMask string, sourceIpAddress string, sourceSubnetMask string) (tr064model.DeleteForwardingEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/layer3forwarding").
		Uri("urn:dslforum-org:service:Layer3Forwarding:1").
		Action("DeleteForwardingEntry").
		AddStringParam("NewDestIPAddress", destIpAddress).
		AddStringParam("NewDestSubnetMask", destSubnetMask).
		AddStringParam("NewSourceIPAddress", sourceIpAddress).
		AddStringParam("NewSourceSubnetMask", sourceSubnetMask).
		Do()
	if err != nil {
		return tr064model.DeleteForwardingEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.DeleteForwardingEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.DeleteForwardingEntryResponse{}, err
	}
	return result, nil
}
