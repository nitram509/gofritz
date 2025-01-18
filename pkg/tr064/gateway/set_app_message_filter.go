package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetAppMessageFilter AUTO-GENERATED (do not edit) code from [x_appsetupSCPD],
// based on SOAP action 'SetAppMessageFilter', Fritz!Box-System-Version 164.08.00
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
func SetAppMessageFilter(session *soap.SoapSession, appId string, aType string, filter string) (tr064model.SetAppMessageFilterResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_appsetup").
		Uri("urn:dslforum-org:service:X_AVM-DE_AppSetup:1").
		Action("SetAppMessageFilter").
		AddStringParam("NewAppId", appId).
		AddStringParam("NewType", aType).
		AddStringParam("NewFilter", filter).
		Do()
	if err != nil {
		return tr064model.SetAppMessageFilterResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetAppMessageFilterResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetAppMessageFilterResponse{}, err
	}
	return result, nil
}
