package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetForwardingEntryEnable AUTO-GENERATED (do not edit) code from [layer3forwardingSCPD],
// based on SOAP action 'SetForwardingEntryEnable', Fritz!Box-System-Version 164.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
func SetForwardingEntryEnable(session *soap.SoapSession, destIpAddress string, destSubnetMask string, sourceIpAddress string, sourceSubnetMask string, enable bool) (tr064model.SetForwardingEntryEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/layer3forwarding").
		Uri("urn:dslforum-org:service:Layer3Forwarding:1").
		Action("SetForwardingEntryEnable").
		AddStringParam("NewDestIPAddress", destIpAddress).
		AddStringParam("NewDestSubnetMask", destSubnetMask).
		AddStringParam("NewSourceIPAddress", sourceIpAddress).
		AddStringParam("NewSourceSubnetMask", sourceSubnetMask).
		AddBoolParam("NewEnable", enable).
		Do().Body.Data
	result := tr064model.SetForwardingEntryEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
