package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2GetWlanConfigurationStatistics AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetStatistics', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2GetWlanConfigurationStatistics(session *soap.SoapSession) (tr064model.GetWlanConfigurationStatisticsResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("GetStatistics").
		Do().Body.Data
	result := tr064model.GetWlanConfigurationStatisticsResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
