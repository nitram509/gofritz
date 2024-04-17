package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAvmSpeedtestInfo AUTO-GENERATED (do not edit) code from [x_speedtestSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_speedtestSCPD]: http://fritz.box:49000/x_speedtestSCPD.xml
func GetAvmSpeedtestInfo(session *soap.SoapSession) (tr064model.GetAvmSpeedtestInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_speedtest").
		Uri("urn:dslforum-org:service:X_AVM-DE_Speedtest:1").
		Action("GetInfo").
		Do().Body.Data
	result := tr064model.GetAvmSpeedtestInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
