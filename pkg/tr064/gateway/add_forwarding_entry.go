package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// AddForwardingEntry AUTO-GENERATED (do not edit) code from [layer3forwardingSCPD],
// based on SOAP action 'AddForwardingEntry', Fritz!Box-System-Version 141.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
func AddForwardingEntry(session *soap.SoapSession, aType string, destIpAddress string, destSubnetMask string, sourceIpAddress string, sourceSubnetMask string, gatewayIpAddress string, aInterface string, forwardingMetric int) (tr064model.AddForwardingEntryResponse, error) {
	bodyData := soap.NewSoapRequest(session).
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
		Do().Body.Data
	result := tr064model.AddForwardingEntryResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
