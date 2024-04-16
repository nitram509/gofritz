package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetDVBCEnable AUTO-GENERATED (do not edit) code from [x_mediaSCPD],
// based on SOAP action 'GetDVBCEnable', Fritz!Box-System-Version 141.07.57
//
// [x_mediaSCPD]: http://fritz.box:49000/x_mediaSCPD.xml
func GetDVBCEnable(session *soap.SoapSession) (tr064model.GetDVBCEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_media").
		Uri("urn:dslforum-org:service:X_AVM-DE_Media:1").
		Action("GetDVBCEnable").
		Do().Body.Data
	result := tr064model.GetDVBCEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
