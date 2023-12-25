package wan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetTotalBytesSent AUTO-GENERATED (do not edit) code from [wancommonifconfigSCPD],
// based on SOAP action 'GetTotalBytesSent', Fritz!Box-System-Version 141.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
func GetTotalBytesSent(session *soap.SoapSession) (tr064model.GetTotalBytesSentResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wancommonifconfig1").
		Uri("urn:dslforum-org:service:WANCommonInterfaceConfig:1").
		Action("GetTotalBytesSent").
		Do().Body.Data
	result := tr064model.GetTotalBytesSentResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
