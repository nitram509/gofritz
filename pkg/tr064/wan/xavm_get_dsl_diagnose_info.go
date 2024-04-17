package wan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmGetDSLDiagnoseInfo AUTO-GENERATED (do not edit) code from [wandslifconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetDSLDiagnoseInfo', Fritz!Box-System-Version 164.07.57
//
// [wandslifconfigSCPD]: http://fritz.box:49000/wandslifconfigSCPD.xml
func XavmGetDSLDiagnoseInfo(session *soap.SoapSession) (tr064model.XavmGetDSLDiagnoseInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wandslifconfig1").
		Uri("urn:dslforum-org:service:WANDSLInterfaceConfig:1").
		Action("X_AVM-DE_GetDSLDiagnoseInfo").
		Do()
	if err != nil {
		return tr064model.XavmGetDSLDiagnoseInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmGetDSLDiagnoseInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmGetDSLDiagnoseInfoResponse{}, err
	}
	return result, nil
}
