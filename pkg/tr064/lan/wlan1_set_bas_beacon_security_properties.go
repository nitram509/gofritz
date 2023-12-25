package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan1SetBasBeaconSecurityProperties AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetBasBeaconSecurityProperties', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1SetBasBeaconSecurityProperties(session *soap.SoapSession, basicEncryptionModes string, basicAuthenticationMode string) (tr064model.SetBasBeaconSecurityPropertiesResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("SetBasBeaconSecurityProperties").
		AddStringParam("NewBasicEncryptionModes", basicEncryptionModes).
		AddStringParam("NewBasicAuthenticationMode", basicAuthenticationMode).
		Do().Body.Data
	result := tr064model.SetBasBeaconSecurityPropertiesResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
