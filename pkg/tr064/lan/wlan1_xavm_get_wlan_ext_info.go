package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1XavmGetWLANExtInfo AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetWLANExtInfo', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1XavmGetWLANExtInfo(session *soap.SoapSession) (tr064model.XavmGetWLANExtInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("X_AVM-DE_GetWLANExtInfo").
		Do()
	if err != nil {
		return tr064model.XavmGetWLANExtInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetWLANExtInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetWLANExtInfoResponse{}, err
	}
	return result, nil
}
