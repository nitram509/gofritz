package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan1GetChannelInfo AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetChannelInfo', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1GetChannelInfo(session *soap.SoapSession) (tr064model.GetChannelInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("GetChannelInfo").
		Do().Body.Data
	result := tr064model.GetChannelInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
