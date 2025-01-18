package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetIPRoutersList AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetIPRoutersList', Fritz!Box-System-Version 164.08.00
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func GetIPRoutersList(session *soap.SoapSession) (tr064model.GetIPRoutersListResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("GetIPRoutersList").
		Do()
	if err != nil {
		return tr064model.GetIPRoutersListResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetIPRoutersListResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetIPRoutersListResponse{}, err
	}
	return result, nil
}
