package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan1GetBeaconType AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetBeaconType', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan1GetBeaconType(session *soap.SoapSession) (tr064model.GetBeaconTypeResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig1").
		Uri("urn:dslforum-org:service:WLANConfiguration:1").
		Action("GetBeaconType").
		Do().Body.Data
	result := tr064model.GetBeaconTypeResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
