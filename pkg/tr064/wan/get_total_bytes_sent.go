package wan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetTotalBytesSent AUTO-GENERATED (do not edit) code from [wancommonifconfigSCPD],
// based on SOAP action 'GetTotalBytesSent', Fritz!Box-System-Version 164.08.00
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
func GetTotalBytesSent(session *soap.SoapSession) (tr064model.GetTotalBytesSentResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wancommonifconfig1").
		Uri("urn:dslforum-org:service:WANCommonInterfaceConfig:1").
		Action("GetTotalBytesSent").
		Do()
	if err != nil {
		return tr064model.GetTotalBytesSentResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetTotalBytesSentResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetTotalBytesSentResponse{}, err
	}
	return result, nil
}
