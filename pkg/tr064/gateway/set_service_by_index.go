package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetServiceByIndex AUTO-GENERATED (do not edit) code from [x_myfritzSCPD],
// based on SOAP action 'SetServiceByIndex', Fritz!Box-System-Version 141.07.57
//
// [x_myfritzSCPD]: http://fritz.box:49000/x_myfritzSCPD.xml
func SetServiceByIndex(session *soap.SoapSession, index int, enabled bool, name string, scheme string, port int, urlPath string, aType string, ipv4Address string, ipv6Address string, ipv6InterfaceId string, macAddress string, hostName string) (tr064model.SetServiceByIndexResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_myfritz").
		Uri("urn:dslforum-org:service:X_AVM-DE_MyFritz:1").
		Action("SetServiceByIndex").
		AddIntParam("NewIndex", index).
		AddBoolParam("NewEnabled", enabled).
		AddStringParam("NewName", name).
		AddStringParam("NewScheme", scheme).
		AddIntParam("NewPort", port).
		AddStringParam("NewURLPath", urlPath).
		AddStringParam("NewType", aType).
		AddStringParam("NewIPv4Address", ipv4Address).
		AddStringParam("NewIPv6Address", ipv6Address).
		AddStringParam("NewIPv6InterfaceID", ipv6InterfaceId).
		AddStringParam("NewMACAddress", macAddress).
		AddStringParam("NewHostName", hostName).
		Do().Body.Data
	result := tr064model.SetServiceByIndexResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
