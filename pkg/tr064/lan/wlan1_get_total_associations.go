package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1GetTotalAssociations AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetTotalAssociations', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1GetTotalAssociations(session *soap.SoapSession) (tr064model.GetTotalAssociationsResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("GetTotalAssociations").
		Do()
	if err != nil {
		return tr064model.GetTotalAssociationsResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetTotalAssociationsResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetTotalAssociationsResponse{}, err
	}
	return result, nil
}
