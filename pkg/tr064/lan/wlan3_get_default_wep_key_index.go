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
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("GetDefaultWEPKeyIndex").
		Do().Body.Data
	result := tr064model.GetDefaultWEPKeyIndexResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
