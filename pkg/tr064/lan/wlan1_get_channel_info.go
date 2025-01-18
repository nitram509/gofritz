package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1GetChannelInfo AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetChannelInfo', Fritz!Box-System-Version 164.08.00
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1GetChannelInfo(session *soap.SoapSession) (tr064model.GetChannelInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("GetChannelInfo").
		Do()
	if err != nil {
		return tr064model.GetChannelInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetChannelInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetChannelInfoResponse{}, err
	}
	return result, nil
}
