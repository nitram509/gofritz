package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetLanEthernetIfCfgStatistics AUTO-GENERATED (do not edit) code from [ethifconfigSCPD],
// based on SOAP action 'GetStatistics', Fritz!Box-System-Version 141.07.57
//
// [ethifconfigSCPD]: http://fritz.box:49000/ethifconfigSCPD.xml
func GetLanEthernetIfCfgStatistics(session *soap.SoapSession) (tr064model.GetLanEthernetIfCfgStatisticsResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanethernetifcfg").
		Uri("urn:dslforum-org:service:LANEthernetInterfaceConfig:1").
		Action("GetStatistics").
		Do().Body.Data
	result := tr064model.GetLanEthernetIfCfgStatisticsResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
