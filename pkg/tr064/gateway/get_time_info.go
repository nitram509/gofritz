package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetTimeInfo AUTO-GENERATED (do not edit) code from [timeSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 141.07.57
//
// [timeSCPD]: http://fritz.box:49000/timeSCPD.xml
func GetTimeInfo(session *soap.SoapSession) (tr064model.GetTimeInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/time").
		Uri("urn:dslforum-org:service:Time:1").
		Action("GetInfo").
		Do().Body.Data
	result := tr064model.GetTimeInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
