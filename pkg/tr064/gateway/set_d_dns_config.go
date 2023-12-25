package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetDDNSConfig AUTO-GENERATED (do not edit) code from [x_remoteSCPD],
// based on SOAP action 'SetDDNSConfig', Fritz!Box-System-Version 141.07.57
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
func SetDDNSConfig(session *soap.SoapSession, enabled bool, providerName string, updateUrl string, domain string, username string, mode string, serverIpv4 string, serverIpv6 string, password string) (tr064model.SetDDNSConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_remote").
		Uri("urn:dslforum-org:service:X_AVM-DE_RemoteAccess:1").
		Action("SetDDNSConfig").
		AddBoolParam("NewEnabled", enabled).
		AddStringParam("NewProviderName", providerName).
		AddStringParam("NewUpdateURL", updateUrl).
		AddStringParam("NewDomain", domain).
		AddStringParam("NewUsername", username).
		AddStringParam("NewMode", mode).
		AddStringParam("NewServerIPv4", serverIpv4).
		AddStringParam("NewServerIPv6", serverIpv6).
		AddStringParam("NewPassword", password).
		Do().Body.Data
	result := tr064model.SetDDNSConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
