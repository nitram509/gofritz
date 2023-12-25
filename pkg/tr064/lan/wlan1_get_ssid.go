package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan1GetSSID AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetSSID', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1GetSSID(session *soap.SoapSession) (tr064model.GetSSIDResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("GetSSID").
		Do().Body.Data
	result := tr064model.GetSSIDResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
