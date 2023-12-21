package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

func XAvmGetSpecificHostEntryByIp(session *soap.SoapSession, ipAddress string) (tr064model.XavmGetSpecificHostEntryByIPResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetSpecificHostEntryByIp").
		AddParam("NewIPAddress", ipAddress).
		Do().Body.Data
	result := tr064model.XavmGetSpecificHostEntryByIPResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}

func XAvmGetHostListPath(session *soap.SoapSession) (tr064model.XavmGetHostListPathResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetHostListPath").
		Do().Body.Data
	result := tr064model.XavmGetHostListPathResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
