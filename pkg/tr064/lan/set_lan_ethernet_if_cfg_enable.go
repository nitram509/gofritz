package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetLanEthernetIfCfgEnable AUTO-GENERATED (do not edit) code from [ethifconfigSCPD],
// based on SOAP action 'SetEnable', Fritz!Box-System-Version 164.07.57
//
// [ethifconfigSCPD]: http://fritz.box:49000/ethifconfigSCPD.xml
func SetLanEthernetIfCfgEnable(session *soap.SoapSession, enable bool) (tr064model.SetLanEthernetIfCfgEnableResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanethernetifcfg").
		Uri("urn:dslforum-org:service:LANEthernetInterfaceConfig:1").
		Action("SetEnable").
		AddBoolParam("NewEnable", enable).
		Do()
	if err != nil {
		return tr064model.SetLanEthernetIfCfgEnableResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetLanEthernetIfCfgEnableResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetLanEthernetIfCfgEnableResponse{}, err
	}
	return result, nil
}
