package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAvmSpeedtestInfo AUTO-GENERATED (do not edit) code from [x_speedtestSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [x_speedtestSCPD]: http://fritz.box:49000/x_speedtestSCPD.xml
func GetAvmSpeedtestInfo(session *soap.SoapSession) (tr064model.GetAvmSpeedtestInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_speedtest").
		Uri("urn:dslforum-org:service:X_AVM-DE_Speedtest:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetAvmSpeedtestInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetAvmSpeedtestInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetAvmSpeedtestInfoResponse{}, err
	}
	return result, nil
}
