package wan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetStatisticsTotal AUTO-GENERATED (do not edit) code from [wandslifconfigSCPD],
// based on SOAP action 'GetStatisticsTotal', Fritz!Box-System-Version 164.07.57
//
// [wandslifconfigSCPD]: http://fritz.box:49000/wandslifconfigSCPD.xml
func GetStatisticsTotal(session *soap.SoapSession) (tr064model.GetStatisticsTotalResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wandslifconfig1").
		Uri("urn:dslforum-org:service:WANDSLInterfaceConfig:1").
		Action("GetStatisticsTotal").
		Do()
	if err != nil {
		return tr064model.GetStatisticsTotalResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetStatisticsTotalResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetStatisticsTotalResponse{}, err
	}
	return result, nil
}
