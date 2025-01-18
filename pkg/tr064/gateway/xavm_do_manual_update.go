package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmDoManualUpdate AUTO-GENERATED (do not edit) code from [userifSCPD],
// based on SOAP action 'X_AVM-DE_DoManualUpdate', Fritz!Box-System-Version 164.08.00
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
func XavmDoManualUpdate(session *soap.SoapSession, avmAllowDowngrade bool, avmDownloadUrl string) (tr064model.XavmDoManualUpdateResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/userif").
		Uri("urn:dslforum-org:service:UserInterface:1").
		Action("X_AVM-DE_DoManualUpdate").
		AddBoolParam("NewX_AVM-DE_AllowDowngrade", avmAllowDowngrade).
		AddStringParam("NewX_AVM-DE_DownloadURL", avmDownloadUrl).
		Do()
	if err != nil {
		return tr064model.XavmDoManualUpdateResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmDoManualUpdateResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmDoManualUpdateResponse{}, err
	}
	return result, nil
}
