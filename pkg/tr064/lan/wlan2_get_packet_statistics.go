package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan2GetPacketStatistics AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetPacketStatistics', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2GetPacketStatistics(session *soap.SoapSession) (tr064model.GetPacketStatisticsResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("GetPacketStatistics").
		Do()
	if err != nil {
		return tr064model.GetPacketStatisticsResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetPacketStatisticsResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetPacketStatisticsResponse{}, err
	}
	return result, nil
}
