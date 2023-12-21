package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmGetTR069FirmwareDownloadEnabled AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'X_AVM-DE_GetTR069FirmwareDownloadEnabled', Fritz!Box-System-Version 164.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func XavmGetTR069FirmwareDownloadEnabled(session *soap.SoapSession) (tr064model.XavmGetTR069FirmwareDownloadEnabledResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("X_AVM-DE_GetTR069FirmwareDownloadEnabled").
		Do().Body.Data
	result := tr064model.XavmGetTR069FirmwareDownloadEnabledResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
