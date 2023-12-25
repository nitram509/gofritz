package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetDectListPath AUTO-GENERATED (do not edit) code from [x_dectSCPD],
// based on SOAP action 'GetDectListPath', Fritz!Box-System-Version 141.07.57
//
// [x_dectSCPD]: http://fritz.box:49000/x_dectSCPD.xml
func GetDectListPath(session *soap.SoapSession) (tr064model.GetDectListPathResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_dect").
		Uri("urn:dslforum-org:service:X_AVM-DE_Dect:1").
		Action("GetDectListPath").
		Do().Body.Data
	result := tr064model.GetDectListPathResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
