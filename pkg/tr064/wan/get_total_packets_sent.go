package wan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetTotalPacketsSent AUTO-GENERATED (do not edit) code from [wancommonifconfigSCPD],
// based on SOAP action 'GetTotalPacketsSent', Fritz!Box-System-Version 164.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
func GetTotalPacketsSent(session *soap.SoapSession) (tr064model.GetTotalPacketsSentResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wancommonifconfig1").
		Uri("urn:dslforum-org:service:WANCommonInterfaceConfig:1").
		Action("GetTotalPacketsSent").
		Do()
	if err != nil {
		return tr064model.GetTotalPacketsSentResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetTotalPacketsSentResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetTotalPacketsSentResponse{}, err
	}
	return result, nil
}
