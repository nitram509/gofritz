package wan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetOnlineMonitor AUTO-GENERATED (do not edit) code from [wancommonifconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetOnlineMonitor', Fritz!Box-System-Version 164.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
func XavmGetOnlineMonitor(session *soap.SoapSession, syncGroupIndex int) (tr064model.XavmGetOnlineMonitorResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wancommonifconfig1").
		Uri("urn:dslforum-org:service:WANCommonInterfaceConfig:1").
		Action("X_AVM-DE_GetOnlineMonitor").
		AddIntParam("NewSyncGroupIndex", syncGroupIndex).
		Do()
	if err != nil {
		return tr064model.XavmGetOnlineMonitorResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetOnlineMonitorResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetOnlineMonitorResponse{}, err
	}
	return result, nil
}
