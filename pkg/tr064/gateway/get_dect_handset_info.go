package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetDECTHandsetInfo AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'GetDECTHandsetInfo', Fritz!Box-System-Version 164.08.00
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func GetDECTHandsetInfo(session *soap.SoapSession, dectId int) (tr064model.GetDECTHandsetInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("GetDECTHandsetInfo").
		AddIntParam("NewDectID", dectId).
		Do()
	if err != nil {
		return tr064model.GetDECTHandsetInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetDECTHandsetInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetDECTHandsetInfoResponse{}, err
	}
	return result, nil
}
