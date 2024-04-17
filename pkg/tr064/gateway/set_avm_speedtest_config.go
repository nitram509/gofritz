package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetAvmSpeedtestConfig AUTO-GENERATED (do not edit) code from [x_speedtestSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.07.57
//
// [x_speedtestSCPD]: http://fritz.box:49000/x_speedtestSCPD.xml
func SetAvmSpeedtestConfig(session *soap.SoapSession, enableTcp bool, enableUdp bool, enableUdpBidirect bool, wanEnableTcp bool, wanEnableUdp bool) (tr064model.SetAvmSpeedtestConfigResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_speedtest").
		Uri("urn:dslforum-org:service:X_AVM-DE_Speedtest:1").
		Action("SetConfig").
		AddBoolParam("NewEnableTcp", enableTcp).
		AddBoolParam("NewEnableUdp", enableUdp).
		AddBoolParam("NewEnableUdpBidirect", enableUdpBidirect).
		AddBoolParam("NewWANEnableTcp", wanEnableTcp).
		AddBoolParam("NewWANEnableUdp", wanEnableUdp).
		Do()
	if err != nil {
		return tr064model.SetAvmSpeedtestConfigResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetAvmSpeedtestConfigResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetAvmSpeedtestConfigResponse{}, err
	}
	return result, nil
}
