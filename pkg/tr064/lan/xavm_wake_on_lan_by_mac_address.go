package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmWakeOnLANByMACAddress AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_WakeOnLANByMACAddress', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmWakeOnLANByMACAddress(session *soap.SoapSession, macAddress string) (tr064model.XavmWakeOnLANByMACAddressResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_WakeOnLANByMACAddress").
		AddStringParam("NewMACAddress", macAddress).
		Do()
	if err != nil {
		return tr064model.XavmWakeOnLANByMACAddressResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmWakeOnLANByMACAddressResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmWakeOnLANByMACAddressResponse{}, err
	}
	return result, nil
}
