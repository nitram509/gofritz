package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAvmHomeautoInfo AUTO-GENERATED (do not edit) code from [x_homeautoSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_homeautoSCPD]: http://fritz.box:49000/x_homeautoSCPD.xml
func GetAvmHomeautoInfo(session *soap.SoapSession) (tr064model.GetAvmHomeautoInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_homeauto").
		Uri("urn:dslforum-org:service:X_AVM-DE_Homeauto:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetAvmHomeautoInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetAvmHomeautoInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetAvmHomeautoInfoResponse{}, err
	}
	return result, nil
}
