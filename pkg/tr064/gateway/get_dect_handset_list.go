package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetDECTHandsetList AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'GetDECTHandsetList', Fritz!Box-System-Version 164.08.00
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func GetDECTHandsetList(session *soap.SoapSession) (tr064model.GetDECTHandsetListResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("GetDECTHandsetList").
		Do()
	if err != nil {
		return tr064model.GetDECTHandsetListResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetDECTHandsetListResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetDECTHandsetListResponse{}, err
	}
	return result, nil
}
