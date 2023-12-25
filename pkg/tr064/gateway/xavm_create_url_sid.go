package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmCreateUrlSID AUTO-GENERATED (do not edit) code from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_CreateUrlSID', Fritz!Box-System-Version 141.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
func XavmCreateUrlSID(session *soap.SoapSession) (tr064model.XavmCreateUrlSIDResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceconfig").
		Uri("urn:dslforum-org:service:DeviceConfig:1").
		Action("X_AVM-DE_CreateUrlSID").
		Do().Body.Data
	result := tr064model.XavmCreateUrlSIDResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
