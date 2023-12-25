package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetAppVPN AUTO-GENERATED (do not edit) code from [x_appsetupSCPD],
// based on SOAP action 'SetAppVPN', Fritz!Box-System-Version 141.07.57
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
func SetAppVPN(session *soap.SoapSession, appId string, ipSecIdentifier string, ipSecPreSharedKey string, ipSecXauthUsername string, ipSecXauthPassword string) (tr064model.SetAppVPNResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_appsetup").
		Uri("urn:dslforum-org:service:X_AVM-DE_AppSetup:1").
		Action("SetAppVPN").
		AddStringParam("NewAppId", appId).
		AddStringParam("NewIPSecIdentifier", ipSecIdentifier).
		AddStringParam("NewIPSecPreSharedKey", ipSecPreSharedKey).
		AddStringParam("NewIPSecXauthUsername", ipSecXauthUsername).
		AddStringParam("NewIPSecXauthPassword", ipSecXauthPassword).
		Do().Body.Data
	result := tr064model.SetAppVPNResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
