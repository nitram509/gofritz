package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetSpecificAssociatedDeviceInfo AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetSpecificAssociatedDeviceInfo', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func GetSpecificAssociatedDeviceInfo(session *soap.SoapSession) (tr064model.GetSpecificAssociatedDeviceInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("GetSpecificAssociatedDeviceInfo").
		Do().Body.Data
	result := tr064model.GetSpecificAssociatedDeviceInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
