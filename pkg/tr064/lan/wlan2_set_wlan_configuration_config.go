package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2SetWlanConfigurationConfig AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2SetWlanConfigurationConfig(session *soap.SoapSession, maxBitRate string, channel int, ssid string, beaconType string, macAddressControlEnabled bool, basicEncryptionModes string, basicAuthenticationMode string) (tr064model.SetWlanConfigurationConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("SetConfig").
		AddStringParam("NewMaxBitRate", maxBitRate).
		AddIntParam("NewChannel", channel).
		AddStringParam("NewSSID", ssid).
		AddStringParam("NewBeaconType", beaconType).
		AddBoolParam("NewMACAddressControlEnabled", macAddressControlEnabled).
		AddStringParam("NewBasicEncryptionModes", basicEncryptionModes).
		AddStringParam("NewBasicAuthenticationMode", basicAuthenticationMode).
		Do().Body.Data
	result := tr064model.SetWlanConfigurationConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
