package wan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetTotalBytesReceived AUTO-GENERATED (do not edit) code from [wancommonifconfigSCPD],
// based on SOAP action 'GetTotalBytesReceived', Fritz!Box-System-Version 164.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
func GetTotalBytesReceived(session *soap.SoapSession) (tr064model.GetTotalBytesReceivedResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wancommonifconfig1").
		Uri("urn:dslforum-org:service:WANCommonInterfaceConfig:1").
		Action("GetTotalBytesReceived").
		Do()
	if err != nil {
		return tr064model.GetTotalBytesReceivedResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetTotalBytesReceivedResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetTotalBytesReceivedResponse{}, err
	}
	return result, nil
}
