package tr064model

import "encoding/xml"

// GetSpecificForwardingEntryResponse AUTO-GENERATED (do not edit) model from [layer3forwardingSCPD],
// based on SOAP action 'GetSpecificForwardingEntry', Fritz!Box-System-Version 141.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
type GetSpecificForwardingEntryResponse struct {
	XMLName          xml.Name // rather for debug purpose
	GatewayIPAddress string   `xml:"NewGatewayIPAddress"` //
	Enable           bool     `xml:"NewEnable"`           // default=0
	Status           string   `xml:"NewStatus"`           //
	Type             string   `xml:"NewType"`             //
	Interface        string   `xml:"NewInterface"`        //
	ForwardingMetric int      `xml:"NewForwardingMetric"` // default=0
}
