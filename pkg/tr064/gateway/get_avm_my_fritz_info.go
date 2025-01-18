package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAvmMyFritzInfo AUTO-GENERATED (do not edit) code from [x_myfritzSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [x_myfritzSCPD]: http://fritz.box:49000/x_myfritzSCPD.xml
func GetAvmMyFritzInfo(session *soap.SoapSession) (tr064model.GetAvmMyFritzInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_myfritz").
		Uri("urn:dslforum-org:service:X_AVM-DE_MyFritz:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetAvmMyFritzInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetAvmMyFritzInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetAvmMyFritzInfoResponse{}, err
	}
	return result, nil
}
