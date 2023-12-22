package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmSetHostNameByMACAddress AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_SetHostNameByMACAddress', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmSetHostNameByMACAddress(session *soap.SoapSession, macAddress string, hostName string) (tr064model.XavmSetHostNameByMACAddressResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_SetHostNameByMACAddress").
		AddStringParam("NewMACAddress", macAddress).
		AddStringParam("NewHostName", hostName).
		Do().Body.Data
	result := tr064model.XavmSetHostNameByMACAddressResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
