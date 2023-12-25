package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetLanEthernetIfCfgInfo AUTO-GENERATED (do not edit) code from [ethifconfigSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 141.07.57
//
// [ethifconfigSCPD]: http://fritz.box:49000/ethifconfigSCPD.xml
func GetLanEthernetIfCfgInfo(session *soap.SoapSession) (tr064model.GetLanEthernetIfCfgInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanethernetifcfg").
		Uri("urn:dslforum-org:service:LANEthernetInterfaceConfig:1").
		Action("GetInfo").
		Do().Body.Data
	result := tr064model.GetLanEthernetIfCfgInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
