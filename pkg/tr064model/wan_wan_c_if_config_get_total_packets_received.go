package tr064model

// GetTotalPacketsReceivedResponse auto generated model from [wancommonifconfigSCPD],
// based on SOAP action 'GetTotalPacketsReceived', Fritz!Box-System-Version 164.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
type GetTotalPacketsReceivedResponse struct {
	TotalPacketsReceived int `xml:"NewTotalPacketsReceived"` // default=0
}
