package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan2GetBasBeaconSecurityProperties AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetBasBeaconSecurityProperties', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2GetBasBeaconSecurityProperties(session *soap.SoapSession) (tr064model.GetBasBeaconSecurityPropertiesResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("GetBasBeaconSecurityProperties").
		Do()
	if err != nil {
		return tr064model.GetBasBeaconSecurityPropertiesResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetBasBeaconSecurityPropertiesResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetBasBeaconSecurityPropertiesResponse{}, err
	}
	return result, nil
}
