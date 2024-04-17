package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetAppVPNwithPFS AUTO-GENERATED (do not edit) code from [x_appsetupSCPD],
// based on SOAP action 'SetAppVPNwithPFS', Fritz!Box-System-Version 164.07.57
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
func SetAppVPNwithPFS(session *soap.SoapSession, appId string, ipSecIdentifier string, ipSecPreSharedKey string, ipSecXauthUsername string, ipSecXauthPassword string) (tr064model.SetAppVPNwithPFSResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_appsetup").
		Uri("urn:dslforum-org:service:X_AVM-DE_AppSetup:1").
		Action("SetAppVPNwithPFS").
		AddStringParam("NewAppId", appId).
		AddStringParam("NewIPSecIdentifier", ipSecIdentifier).
		AddStringParam("NewIPSecPreSharedKey", ipSecPreSharedKey).
		AddStringParam("NewIPSecXauthUsername", ipSecXauthUsername).
		AddStringParam("NewIPSecXauthPassword", ipSecXauthPassword).
		Do().Body.Data
	result := tr064model.SetAppVPNwithPFSResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
