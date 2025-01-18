package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetTimeInfo AUTO-GENERATED (do not edit) code from [timeSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [timeSCPD]: http://fritz.box:49000/timeSCPD.xml
func GetTimeInfo(session *soap.SoapSession) (tr064model.GetTimeInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/time").
		Uri("urn:dslforum-org:service:Time:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetTimeInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetTimeInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetTimeInfoResponse{}, err
	}
	return result, nil
}
