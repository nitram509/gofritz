package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetSpecificHostEntry AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'GetSpecificHostEntry', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func GetSpecificHostEntry(session *soap.SoapSession) (tr064model.GetSpecificHostEntryResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("GetSpecificHostEntry").
		Do().Body.Data
	result := tr064model.GetSpecificHostEntryResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
