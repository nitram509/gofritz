package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan3SetDefaultWEPKeyIndex AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetDefaultWEPKeyIndex', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan3SetDefaultWEPKeyIndex(session *soap.SoapSession, defaultWepKeyIndex int) (tr064model.SetDefaultWEPKeyIndexResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("SetDefaultWEPKeyIndex").
		AddIntParam("NewDefaultWEPKeyIndex", defaultWepKeyIndex).
		Do()
	if err != nil {
		return tr064model.SetDefaultWEPKeyIndexResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetDefaultWEPKeyIndexResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetDefaultWEPKeyIndexResponse{}, err
	}
	return result, nil
}
