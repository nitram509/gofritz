package wan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmGetDSLInfo AUTO-GENERATED (do not edit) code from [wandslifconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetDSLInfo', Fritz!Box-System-Version 141.07.57
//
// [wandslifconfigSCPD]: http://fritz.box:49000/wandslifconfigSCPD.xml
func XavmGetDSLInfo(session *soap.SoapSession) (tr064model.XavmGetDSLInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wandslifconfig1").
		Uri("urn:dslforum-org:service:WANDSLInterfaceConfig:1").
		Action("X_AVM-DE_GetDSLInfo").
		Do().Body.Data
	result := tr064model.XavmGetDSLInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
