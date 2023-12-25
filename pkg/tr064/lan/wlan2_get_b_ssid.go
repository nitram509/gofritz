package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2GetBSSID AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetBSSID', Fritz!Box-System-Version 141.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2GetBSSID(session *soap.SoapSession) (tr064model.GetBSSIDResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("GetBSSID").
		Do().Body.Data
	result := tr064model.GetBSSIDResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
