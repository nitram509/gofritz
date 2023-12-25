package wan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetStatisticsTotal AUTO-GENERATED (do not edit) code from [wandslifconfigSCPD],
// based on SOAP action 'GetStatisticsTotal', Fritz!Box-System-Version 141.07.57
//
// [wandslifconfigSCPD]: http://fritz.box:49000/wandslifconfigSCPD.xml
func GetStatisticsTotal(session *soap.SoapSession) (tr064model.GetStatisticsTotalResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wandslifconfig1").
		Uri("urn:dslforum-org:service:WANDSLInterfaceConfig:1").
		Action("GetStatisticsTotal").
		Do().Body.Data
	result := tr064model.GetStatisticsTotalResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
