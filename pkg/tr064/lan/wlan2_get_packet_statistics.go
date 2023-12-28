package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2GetPacketStatistics AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetPacketStatistics', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2GetPacketStatistics(session *soap.SoapSession) (tr064model.GetPacketStatisticsResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("GetPacketStatistics").
		Do().Body.Data
	result := tr064model.GetPacketStatisticsResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
