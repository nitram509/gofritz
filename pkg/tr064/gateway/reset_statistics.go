package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// ResetStatistics AUTO-GENERATED (do not edit) code from [x_speedtestSCPD],
// based on SOAP action 'ResetStatistics', Fritz!Box-System-Version 141.07.57
//
// [x_speedtestSCPD]: http://fritz.box:49000/x_speedtestSCPD.xml
func ResetStatistics(session *soap.SoapSession) (tr064model.ResetStatisticsResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_speedtest").
		Uri("urn:dslforum-org:service:X_AVM-DE_Speedtest:1").
		Action("ResetStatistics").
		Do().Body.Data
	result := tr064model.ResetStatisticsResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
