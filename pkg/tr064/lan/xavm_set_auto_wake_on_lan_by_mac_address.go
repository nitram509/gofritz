package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmSetAutoWakeOnLANByMACAddress AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_SetAutoWakeOnLANByMACAddress', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmSetAutoWakeOnLANByMACAddress(session *soap.SoapSession, macAddress string, autoWolEnabled bool) (tr064model.XavmSetAutoWakeOnLANByMACAddressResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_SetAutoWakeOnLANByMACAddress").
		AddStringParam("NewMACAddress", macAddress).
		AddBoolParam("NewAutoWOLEnabled", autoWolEnabled).
		Do().Body.Data
	result := tr064model.XavmSetAutoWakeOnLANByMACAddressResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
