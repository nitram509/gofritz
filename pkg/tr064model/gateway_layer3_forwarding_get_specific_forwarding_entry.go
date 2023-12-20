package tr064model

// GetSpecificForwardingEntryResponse auto generated model from [layer3forwardingSCPD],
// based on SOAP action 'GetSpecificForwardingEntry', Fritz!Box-System-Version 164.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
type GetSpecificForwardingEntryResponse struct {
	GatewayIPAddress string `xml:"NewGatewayIPAddress"` //
	Enable           bool   `xml:"NewEnable"`           // default=0
	Status           string `xml:"NewStatus"`           //
	Type             string `xml:"NewType"`             //
	Interface        string `xml:"NewInterface"`        //
	ForwardingMetric int    `xml:"NewForwardingMetric"` // default=0
}
