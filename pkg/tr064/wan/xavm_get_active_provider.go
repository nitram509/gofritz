package wan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetActiveProvider AUTO-GENERATED (do not edit) code from [wancommonifconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetActiveProvider', Fritz!Box-System-Version 164.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
func XavmGetActiveProvider(session *soap.SoapSession) (tr064model.XavmGetActiveProviderResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wancommonifconfig1").
		Uri("urn:dslforum-org:service:WANCommonInterfaceConfig:1").
		Action("X_AVM-DE_GetActiveProvider").
		Do()
	if err != nil {
		return tr064model.XavmGetActiveProviderResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetActiveProviderResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetActiveProviderResponse{}, err
	}
	return result, nil
}
