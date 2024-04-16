package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetSearchProgress AUTO-GENERATED (do not edit) code from [x_mediaSCPD],
// based on SOAP action 'GetSearchProgress', Fritz!Box-System-Version 141.07.57
//
// [x_mediaSCPD]: http://fritz.box:49000/x_mediaSCPD.xml
func GetSearchProgress(session *soap.SoapSession) (tr064model.GetSearchProgressResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_media").
		Uri("urn:dslforum-org:service:X_AVM-DE_Media:1").
		Action("GetSearchProgress").
		Do().Body.Data
	result := tr064model.GetSearchProgressResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
