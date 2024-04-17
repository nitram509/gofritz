package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetLanHCfgMgmInfo AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func GetLanHCfgMgmInfo(session *soap.SoapSession) (tr064model.GetLanHCfgMgmInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetLanHCfgMgmInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetLanHCfgMgmInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetLanHCfgMgmInfoResponse{}, err
	}
	return result, nil
}
