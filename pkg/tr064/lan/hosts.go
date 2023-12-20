package lan

import (
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

func XAvmGetSpecificHostEntryByIp(session soap.SoapSession, ipAddress string) tr064model.XAvmGetSpecificHostEntryByIpResponse {
	return soap.NewSoapRequest().
		WithAuthenticator(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetSpecificHostEntryByIp").
		AddParam("NewIPAddress", ipAddress).
		Do().Body.XAvmGetSpecificHostEntryByIpResponse
}

func XAvmGetHostListPath(session soap.SoapSession) tr064model.XAvmGetHostListPathResponse {
	return soap.NewSoapRequest().
		WithAuthenticator(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetHostListPath").
		Do().Body.XAvmGetHostListPathResponse
}
