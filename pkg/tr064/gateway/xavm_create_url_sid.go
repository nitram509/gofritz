package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmCreateUrlSID AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_CreateUrlSID', Fritz!Box-System-Version 164.08.00
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func XavmCreateUrlSID(session *soap.SoapSession) (tr064model.XavmCreateUrlSIDResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("X_AVM-DE_CreateUrlSID").
		Do()
	if err != nil {
		return tr064model.XavmCreateUrlSIDResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmCreateUrlSIDResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmCreateUrlSIDResponse{}, err
	}
	return result, nil
}
