package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetAvmSpeedtestStatistics AUTO-GENERATED (do not edit) code from [x_speedtestSCPD],
// based on SOAP action 'GetStatistics', Fritz!Box-System-Version 141.07.57
//
// [x_speedtestSCPD]: http://fritz.box:49000/x_speedtestSCPD.xml
func GetAvmSpeedtestStatistics(session *soap.SoapSession) (tr064model.GetAvmSpeedtestStatisticsResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_speedtest").
		Uri("urn:dslforum-org:service:X_AVM-DE_Speedtest:1").
		Action("GetStatistics").
		Do().Body.Data
	result := tr064model.GetAvmSpeedtestStatisticsResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
