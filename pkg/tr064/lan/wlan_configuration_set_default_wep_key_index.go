package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetDefaultWEPKeyIndex AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'SetDefaultWEPKeyIndex', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func SetDefaultWEPKeyIndex(session *soap.SoapSession, defaultWepKeyIndex int) (tr064model.SetDefaultWEPKeyIndexResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("SetDefaultWEPKeyIndex").
		AddIntParam("NewDefaultWEPKeyIndex", defaultWepKeyIndex).
		Do().Body.Data
	result := tr064model.SetDefaultWEPKeyIndexResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
