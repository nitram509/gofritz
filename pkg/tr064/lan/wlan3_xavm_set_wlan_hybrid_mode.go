package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan3XavmSetWLANHybridMode AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetWLANHybridMode', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan3XavmSetWLANHybridMode(session *soap.SoapSession, enable bool, beaconType string, keyPassphrase string, ssid string, bSsid string, trafficMode string, manualSpeed bool, maxSpeedDS int, maxSpeedUS int) (tr064model.XavmSetWLANHybridModeResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("X_AVM-DE_SetWLANHybridMode").
		AddBoolParam("NewEnable", enable).
		AddStringParam("NewBeaconType", beaconType).
		AddStringParam("NewKeyPassphrase", keyPassphrase).
		AddStringParam("NewSSID", ssid).
		AddStringParam("NewBSSID", bSsid).
		AddStringParam("NewTrafficMode", trafficMode).
		AddBoolParam("NewManualSpeed", manualSpeed).
		AddIntParam("NewMaxSpeedDS", maxSpeedDS).
		AddIntParam("NewMaxSpeedUS", maxSpeedUS).
		Do().Body.Data
	result := tr064model.XavmSetWLANHybridModeResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
