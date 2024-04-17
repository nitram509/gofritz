package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// Wlan1SetChannel AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetChannel', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1SetChannel(session *soap.SoapSession, channel int) (tr064model.SetChannelResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("SetChannel").
		AddIntParam("NewChannel", channel).
		Do()
	if err != nil {
		return tr064model.SetChannelResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetChannelResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetChannelResponse{}, err
	}
	return result, nil
}
