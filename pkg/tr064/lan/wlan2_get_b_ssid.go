package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan2GetBSSID AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetBSSID', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2GetBSSID(session *soap.SoapSession) (tr064model.GetBSSIDResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("GetBSSID").
		Do()
	if err != nil {
		return tr064model.GetBSSIDResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetBSSIDResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetBSSIDResponse{}, err
	}
	return result, nil
}
