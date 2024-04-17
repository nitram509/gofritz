package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetDectListPath AUTO-GENERATED (do not edit) code from [x_dectSCPD],
// based on SOAP action 'GetDectListPath', Fritz!Box-System-Version 164.07.57
//
// [x_dectSCPD]: http://fritz.box:49000/x_dectSCPD.xml
func GetDectListPath(session *soap.SoapSession) (tr064model.GetDectListPathResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_dect").
		Uri("urn:dslforum-org:service:X_AVM-DE_Dect:1").
		Action("GetDectListPath").
		Do()
	if err != nil {
		return tr064model.GetDectListPathResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetDectListPathResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetDectListPathResponse{}, err
	}
	return result, nil
}
