package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetUserInfo AUTO-GENERATED (do not edit) code from [x_storageSCPD],
// based on SOAP action 'GetUserInfo', Fritz!Box-System-Version 164.07.57
//
// [x_storageSCPD]: http://fritz.box:49000/x_storageSCPD.xml
func GetUserInfo(session *soap.SoapSession) (tr064model.GetUserInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_storage").
		Uri("urn:dslforum-org:service:X_AVM-DE_Storage:1").
		Action("GetUserInfo").
		Do()
	if err != nil {
		return tr064model.GetUserInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetUserInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetUserInfoResponse{}, err
	}
	return result, nil
}
