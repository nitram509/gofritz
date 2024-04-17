package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan3GetDefaultWEPKeyIndex AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetDefaultWEPKeyIndex', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan3GetDefaultWEPKeyIndex(session *soap.SoapSession) (tr064model.GetDefaultWEPKeyIndexResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("GetDefaultWEPKeyIndex").
		Do()
	if err != nil {
		return tr064model.GetDefaultWEPKeyIndexResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetDefaultWEPKeyIndexResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetDefaultWEPKeyIndexResponse{}, err
	}
	return result, nil
}
