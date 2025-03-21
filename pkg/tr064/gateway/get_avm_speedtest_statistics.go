package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAvmSpeedtestStatistics AUTO-GENERATED (do not edit) code from [x_speedtestSCPD],
// based on SOAP action 'GetStatistics', Fritz!Box-System-Version 164.08.00
//
// [x_speedtestSCPD]: http://fritz.box:49000/x_speedtestSCPD.xml
func GetAvmSpeedtestStatistics(session *soap.SoapSession) (tr064model.GetAvmSpeedtestStatisticsResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_speedtest").
		Uri("urn:dslforum-org:service:X_AVM-DE_Speedtest:1").
		Action("GetStatistics").
		Do()
	if err != nil {
		return tr064model.GetAvmSpeedtestStatisticsResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetAvmSpeedtestStatisticsResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetAvmSpeedtestStatisticsResponse{}, err
	}
	return result, nil
}
