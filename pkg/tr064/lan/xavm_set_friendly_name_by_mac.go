package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetFriendlyNameByMAC AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_SetFriendlyNameByMAC', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmSetFriendlyNameByMAC(session *soap.SoapSession, macAddress string, avmFriendlyName string) (tr064model.XavmSetFriendlyNameByMACResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_SetFriendlyNameByMAC").
		AddStringParam("NewMACAddress", macAddress).
		AddStringParam("NewX_AVM-DE_FriendlyName", avmFriendlyName).
		Do()
	if err != nil {
		return tr064model.XavmSetFriendlyNameByMACResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetFriendlyNameByMACResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetFriendlyNameByMACResponse{}, err
	}
	return result, nil
}
