package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan2SetBasBeaconSecurityProperties AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetBasBeaconSecurityProperties', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2SetBasBeaconSecurityProperties(session *soap.SoapSession, basicEncryptionModes string, basicAuthenticationMode string) (tr064model.SetBasBeaconSecurityPropertiesResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("SetBasBeaconSecurityProperties").
		AddStringParam("NewBasicEncryptionModes", basicEncryptionModes).
		AddStringParam("NewBasicAuthenticationMode", basicAuthenticationMode).
		Do()
	if err != nil {
		return tr064model.SetBasBeaconSecurityPropertiesResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetBasBeaconSecurityPropertiesResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetBasBeaconSecurityPropertiesResponse{}, err
	}
	return result, nil
}
