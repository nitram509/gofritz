package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetHostListPath AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetHostListPath', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmGetHostListPath(session *soap.SoapSession) (tr064model.XavmGetHostListPathResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetHostListPath").
		Do().Body.Data
	result := tr064model.XavmGetHostListPathResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
