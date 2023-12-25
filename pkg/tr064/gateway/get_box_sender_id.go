package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetBoxSenderId AUTO-GENERATED (do not edit) code from [x_appsetupSCPD],
// based on SOAP action 'GetBoxSenderId', Fritz!Box-System-Version 141.07.57
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
func GetBoxSenderId(session *soap.SoapSession, appId string) (tr064model.GetBoxSenderIdResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_appsetup").
		Uri("urn:dslforum-org:service:X_AVM-DE_AppSetup:1").
		Action("GetBoxSenderId").
		AddStringParam("NewAppId", appId).
		Do().Body.Data
	result := tr064model.GetBoxSenderIdResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
