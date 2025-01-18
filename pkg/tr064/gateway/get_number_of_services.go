package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetNumberOfServices AUTO-GENERATED (do not edit) code from [x_myfritzSCPD],
// based on SOAP action 'GetNumberOfServices', Fritz!Box-System-Version 164.08.00
//
// [x_myfritzSCPD]: http://fritz.box:49000/x_myfritzSCPD.xml
func GetNumberOfServices(session *soap.SoapSession) (tr064model.GetNumberOfServicesResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_myfritz").
		Uri("urn:dslforum-org:service:X_AVM-DE_MyFritz:1").
		Action("GetNumberOfServices").
		Do()
	if err != nil {
		return tr064model.GetNumberOfServicesResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetNumberOfServicesResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetNumberOfServicesResponse{}, err
	}
	return result, nil
}
