package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetLanEthernetIfCfgInfo AUTO-GENERATED (do not edit) code from [ethifconfigSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [ethifconfigSCPD]: http://fritz.box:49000/ethifconfigSCPD.xml
func GetLanEthernetIfCfgInfo(session *soap.SoapSession) (tr064model.GetLanEthernetIfCfgInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanethernetifcfg").
		Uri("urn:dslforum-org:service:LANEthernetInterfaceConfig:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetLanEthernetIfCfgInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetLanEthernetIfCfgInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetLanEthernetIfCfgInfoResponse{}, err
	}
	return result, nil
}
