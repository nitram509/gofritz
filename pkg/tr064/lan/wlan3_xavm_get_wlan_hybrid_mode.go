package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan3XavmGetWLANHybridMode AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetWLANHybridMode', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan3XavmGetWLANHybridMode(session *soap.SoapSession) (tr064model.XavmGetWLANHybridModeResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("X_AVM-DE_GetWLANHybridMode").
		Do()
	if err != nil {
		return tr064model.XavmGetWLANHybridModeResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetWLANHybridModeResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetWLANHybridModeResponse{}, err
	}
	return result, nil
}
