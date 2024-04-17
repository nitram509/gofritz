package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// AddForwardingEntry AUTO-GENERATED (do not edit) code from [layer3forwardingSCPD],
// based on SOAP action 'AddForwardingEntry', Fritz!Box-System-Version 164.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
func AddForwardingEntry(session *soap.SoapSession, aType string, destIpAddress string, destSubnetMask string, sourceIpAddress string, sourceSubnetMask string, gatewayIpAddress string, aInterface string, forwardingMetric int) (tr064model.AddForwardingEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/layer3forwarding").
		Uri("urn:dslforum-org:service:Layer3Forwarding:1").
		Action("AddForwardingEntry").
		AddStringParam("NewType", aType).
		AddStringParam("NewDestIPAddress", destIpAddress).
		AddStringParam("NewDestSubnetMask", destSubnetMask).
		AddStringParam("NewSourceIPAddress", sourceIpAddress).
		AddStringParam("NewSourceSubnetMask", sourceSubnetMask).
		AddStringParam("NewGatewayIPAddress", gatewayIpAddress).
		AddStringParam("NewInterface", aInterface).
		AddIntParam("NewForwardingMetric", forwardingMetric).
		Do()
	if err != nil {
		return tr064model.AddForwardingEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.AddForwardingEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.AddForwardingEntryResponse{}, err
	}
	return result, nil
}
