package lan

import (
	"github.com/nitram509/gofitz/pkg/models"
	"github.com/nitram509/gofitz/pkg/soap"
)

func XAvmGetSpecificHostEntryByIp(session soap.SoapSession, ipAddress string) models.XAvmGetSpecificHostEntryByIpResponse {
	return soap.NewSoapRequest().
		WithAuthenticator(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetSpecificHostEntryByIp").
		AddParam("NewIPAddress", ipAddress).
		Do().Body.XAvmGetSpecificHostEntryByIpResponse
}

func XAvmGetHostListPath(session soap.SoapSession) models.XAvmGetHostListPathResponse {
	return soap.NewSoapRequest().
		WithAuthenticator(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetHostListPath").
		Do().Body.XAvmGetHostListPathResponse
}
