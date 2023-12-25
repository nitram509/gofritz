package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetGenericHostEntry AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'GetGenericHostEntry', Fritz!Box-System-Version 141.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func GetGenericHostEntry(session *soap.SoapSession, index int) (tr064model.GetGenericHostEntryResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("GetGenericHostEntry").
		AddIntParam("NewIndex", index).
		Do().Body.Data
	result := tr064model.GetGenericHostEntryResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
