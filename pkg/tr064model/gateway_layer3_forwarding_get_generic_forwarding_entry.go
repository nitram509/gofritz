package tr064model

// GetGenericForwardingEntryResponse auto generated model from [layer3forwardingSCPD],
// based on SOAP action 'GetGenericForwardingEntry', Fritz!Box-System-Version 164.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
type GetGenericForwardingEntryResponse struct {
	Enable           bool   `xml:"NewEnable"`           // default=0
	Status           string `xml:"NewStatus"`           //
	Type             string `xml:"NewType"`             //
	DestIPAddress    string `xml:"NewDestIPAddress"`    //
	DestSubnetMask   string `xml:"NewDestSubnetMask"`   //
	SourceIPAddress  string `xml:"NewSourceIPAddress"`  //
	SourceSubnetMask string `xml:"NewSourceSubnetMask"` //
	GatewayIPAddress string `xml:"NewGatewayIPAddress"` //
	Interface        string `xml:"NewInterface"`        //
	ForwardingMetric int    `xml:"NewForwardingMetric"` // default=0
}
