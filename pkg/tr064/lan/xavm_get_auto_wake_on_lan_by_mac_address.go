package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetAutoWakeOnLANByMACAddress AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetAutoWakeOnLANByMACAddress', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmGetAutoWakeOnLANByMACAddress(session *soap.SoapSession, macAddress string) (tr064model.XavmGetAutoWakeOnLANByMACAddressResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetAutoWakeOnLANByMACAddress").
		AddStringParam("NewMACAddress", macAddress).
		Do()
	if err != nil {
		return tr064model.XavmGetAutoWakeOnLANByMACAddressResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetAutoWakeOnLANByMACAddressResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetAutoWakeOnLANByMACAddressResponse{}, err
	}
	return result, nil
}
