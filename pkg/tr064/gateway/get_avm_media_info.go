package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAvmMediaInfo AUTO-GENERATED (do not edit) code from [x_mediaSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 141.07.57
//
// [x_mediaSCPD]: http://fritz.box:49000/x_mediaSCPD.xml
func GetAvmMediaInfo(session *soap.SoapSession) (tr064model.GetAvmMediaInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_media").
		Uri("urn:dslforum-org:service:X_AVM-DE_Media:1").
		Action("GetInfo").
		Do().Body.Data
	result := tr064model.GetAvmMediaInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
