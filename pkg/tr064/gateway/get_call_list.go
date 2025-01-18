package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetCallList AUTO-GENERATED (do not edit) code from [x_contactSCPD],
// based on SOAP action 'GetCallList', Fritz!Box-System-Version 164.08.00
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
func GetCallList(session *soap.SoapSession) (tr064model.GetCallListResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_contact").
		Uri("urn:dslforum-org:service:X_AVM-DE_OnTel:1").
		Action("GetCallList").
		Do()
	if err != nil {
		return tr064model.GetCallListResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetCallListResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetCallListResponse{}, err
	}
	return result, nil
}
