package tr064model

// GetTotalBytesReceivedResponse auto generated model from [wancommonifconfigSCPD],
// based on SOAP action 'GetTotalBytesReceived', Fritz!Box-System-Version 164.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
type GetTotalBytesReceivedResponse struct {
	TotalBytesReceived int `xml:"NewTotalBytesReceived"` // default=0
}
