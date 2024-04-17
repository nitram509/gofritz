package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetTR069FirmwareDownloadEnabled AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'X_AVM-DE_SetTR069FirmwareDownloadEnabled', Fritz!Box-System-Version 164.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func XavmSetTR069FirmwareDownloadEnabled(session *soap.SoapSession, tr069FirmwareDownloadEnabled bool) (tr064model.XavmSetTR069FirmwareDownloadEnabledResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("X_AVM-DE_SetTR069FirmwareDownloadEnabled").
		AddBoolParam("NewTR069FirmwareDownloadEnabled", tr069FirmwareDownloadEnabled).
		Do()
	if err != nil {
		return tr064model.XavmSetTR069FirmwareDownloadEnabledResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetTR069FirmwareDownloadEnabledResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetTR069FirmwareDownloadEnabledResponse{}, err
	}
	return result, nil
}
