package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetAvmSpeedtestConfig AUTO-GENERATED (do not edit) code from [x_speedtestSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.07.57
//
// [x_speedtestSCPD]: http://fritz.box:49000/x_speedtestSCPD.xml
func SetAvmSpeedtestConfig(session *soap.SoapSession) (tr064model.SetAvmSpeedtestConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_speedtest").
		Uri("urn:dslforum-org:service:X_AVM-DE_Speedtest:1").
		Action("SetConfig").
		Do().Body.Data
	result := tr064model.SetAvmSpeedtestConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
