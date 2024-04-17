package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetSecurityPort AUTO-GENERATED (do not edit) code from [deviceinfoSCPD],
// based on SOAP action 'GetSecurityPort', Fritz!Box-System-Version 164.07.57
//
// [deviceinfoSCPD]: http://fritz.box:49000/deviceinfoSCPD.xml
func GetSecurityPort(session *soap.SoapSession) (tr064model.GetSecurityPortResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceinfo").
		Uri("urn:dslforum-org:service:DeviceInfo:1").
		Action("GetSecurityPort").
		Do()
	if err != nil {
		return tr064model.GetSecurityPortResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetSecurityPortResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetSecurityPortResponse{}, err
	}
	return result, nil
}
