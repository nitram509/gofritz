package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan2GetWlanConfigurationStatistics AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetStatistics', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2GetWlanConfigurationStatistics(session *soap.SoapSession) (tr064model.GetWlanConfigurationStatisticsResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("GetStatistics").
		Do()
	if err != nil {
		return tr064model.GetWlanConfigurationStatisticsResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetWlanConfigurationStatisticsResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetWlanConfigurationStatisticsResponse{}, err
	}
	return result, nil
}
