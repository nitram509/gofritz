package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetList AUTO-GENERATED (do not edit) code from [x_tamSCPD],
// based on SOAP action 'GetList', Fritz!Box-System-Version 164.08.00
//
// [x_tamSCPD]: http://fritz.box:49000/x_tamSCPD.xml
func GetList(session *soap.SoapSession) (tr064model.GetListResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_tam").
		Uri("urn:dslforum-org:service:X_AVM-DE_TAM:1").
		Action("GetList").
		Do()
	if err != nil {
		return tr064model.GetListResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetListResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetListResponse{}, err
	}
	return result, nil
}
