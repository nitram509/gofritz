package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetAutoWakeOnLANByMACAddress AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_SetAutoWakeOnLANByMACAddress', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmSetAutoWakeOnLANByMACAddress(session *soap.SoapSession, macAddress string, autoWolEnabled bool) (tr064model.XavmSetAutoWakeOnLANByMACAddressResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_SetAutoWakeOnLANByMACAddress").
		AddStringParam("NewMACAddress", macAddress).
		AddBoolParam("NewAutoWOLEnabled", autoWolEnabled).
		Do()
	if err != nil {
		return tr064model.XavmSetAutoWakeOnLANByMACAddressResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmSetAutoWakeOnLANByMACAddressResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmSetAutoWakeOnLANByMACAddressResponse{}, err
	}
	return result, nil
}
