package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2SetChannel AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetChannel', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2SetChannel(session *soap.SoapSession, channel int) (tr064model.SetChannelResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("SetChannel").
		AddIntParam("NewChannel", channel).
		Do().Body.Data
	result := tr064model.SetChannelResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
