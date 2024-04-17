package wan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetTotalPacketsReceived AUTO-GENERATED (do not edit) code from [wancommonifconfigSCPD],
// based on SOAP action 'GetTotalPacketsReceived', Fritz!Box-System-Version 164.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
func GetTotalPacketsReceived(session *soap.SoapSession) (tr064model.GetTotalPacketsReceivedResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wancommonifconfig1").
		Uri("urn:dslforum-org:service:WANCommonInterfaceConfig:1").
		Action("GetTotalPacketsReceived").
		Do()
	if err != nil {
		return tr064model.GetTotalPacketsReceivedResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetTotalPacketsReceivedResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetTotalPacketsReceivedResponse{}, err
	}
	return result, nil
}
