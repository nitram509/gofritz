package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmWakeOnLANByMACAddress AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_WakeOnLANByMACAddress', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmWakeOnLANByMACAddress(session *soap.SoapSession, macAddress string) (tr064model.XavmWakeOnLANByMACAddressResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_WakeOnLANByMACAddress").
		AddStringParam("NewMACAddress", macAddress).
		Do().Body.Data
	result := tr064model.XavmWakeOnLANByMACAddressResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
