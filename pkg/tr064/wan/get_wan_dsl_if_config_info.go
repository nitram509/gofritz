package wan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetWanDslIfConfigInfo AUTO-GENERATED (do not edit) code from [wandslifconfigSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [wandslifconfigSCPD]: http://fritz.box:49000/wandslifconfigSCPD.xml
func GetWanDslIfConfigInfo(session *soap.SoapSession) (tr064model.GetWanDslIfConfigInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wandslifconfig1").
		Uri("urn:dslforum-org:service:WANDSLInterfaceConfig:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetWanDslIfConfigInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetWanDslIfConfigInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetWanDslIfConfigInfoResponse{}, err
	}
	return result, nil
}
