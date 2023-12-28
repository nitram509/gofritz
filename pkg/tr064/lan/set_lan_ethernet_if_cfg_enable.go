package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetLanEthernetIfCfgEnable AUTO-GENERATED (do not edit) code from [ethifconfigSCPD],
// based on SOAP action 'SetEnable', Fritz!Box-System-Version 164.07.57
//
// [ethifconfigSCPD]: http://fritz.box:49000/ethifconfigSCPD.xml
func SetLanEthernetIfCfgEnable(session *soap.SoapSession, enable bool) (tr064model.SetLanEthernetIfCfgEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanethernetifcfg").
		Uri("urn:dslforum-org:service:LANEthernetInterfaceConfig:1").
		Action("SetEnable").
		AddBoolParam("NewEnable", enable).
		Do().Body.Data
	result := tr064model.SetLanEthernetIfCfgEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
