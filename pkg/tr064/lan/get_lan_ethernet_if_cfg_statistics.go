package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetLanEthernetIfCfgStatistics AUTO-GENERATED (do not edit) code from [ethifconfigSCPD],
// based on SOAP action 'GetStatistics', Fritz!Box-System-Version 164.07.57
//
// [ethifconfigSCPD]: http://fritz.box:49000/ethifconfigSCPD.xml
func GetLanEthernetIfCfgStatistics(session *soap.SoapSession) (tr064model.GetLanEthernetIfCfgStatisticsResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanethernetifcfg").
		Uri("urn:dslforum-org:service:LANEthernetInterfaceConfig:1").
		Action("GetStatistics").
		Do()
	if err != nil {
		return tr064model.GetLanEthernetIfCfgStatisticsResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetLanEthernetIfCfgStatisticsResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetLanEthernetIfCfgStatisticsResponse{}, err
	}
	return result, nil
}
