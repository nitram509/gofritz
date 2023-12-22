package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmDoManualUpdate AUTO-GENERATED (do not edit) code from [userifSCPD],
// based on SOAP action 'X_AVM-DE_DoManualUpdate', Fritz!Box-System-Version 164.07.57
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
func XavmDoManualUpdate(session *soap.SoapSession, avmAllowDowngrade bool, avmDownloadUrl string) (tr064model.XavmDoManualUpdateResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/userif").
		Uri("urn:dslforum-org:service:UserInterface:1").
		Action("X_AVM-DE_DoManualUpdate").
		AddBoolParam("NewX_AVM-DE_AllowDowngrade", avmAllowDowngrade).
		AddStringParam("NewX_AVM-DE_DownloadURL", avmDownloadUrl).
		Do().Body.Data
	result := tr064model.XavmDoManualUpdateResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
