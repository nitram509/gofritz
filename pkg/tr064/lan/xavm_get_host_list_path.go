package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetHostListPath AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetHostListPath', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmGetHostListPath(session *soap.SoapSession) (tr064model.XavmGetHostListPathResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetHostListPath").
		Do()
	if err != nil {
		return tr064model.XavmGetHostListPathResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetHostListPathResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetHostListPathResponse{}, err
	}
	return result, nil
}
