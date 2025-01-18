package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// DectDoUpdate AUTO-GENERATED (do not edit) code from [x_dectSCPD],
// based on SOAP action 'DectDoUpdate', Fritz!Box-System-Version 164.08.00
//
// [x_dectSCPD]: http://fritz.box:49000/x_dectSCPD.xml
func DectDoUpdate(session *soap.SoapSession, id string) (tr064model.DectDoUpdateResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_dect").
		Uri("urn:dslforum-org:service:X_AVM-DE_Dect:1").
		Action("DectDoUpdate").
		AddStringParam("NewID", id).
		Do()
	if err != nil {
		return tr064model.DectDoUpdateResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.DectDoUpdateResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.DectDoUpdateResponse{}, err
	}
	return result, nil
}
