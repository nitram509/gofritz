package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetSMBServer AUTO-GENERATED (do not edit) code from [x_storageSCPD],
// based on SOAP action 'SetSMBServer', Fritz!Box-System-Version 164.07.57
//
// [x_storageSCPD]: http://fritz.box:49000/x_storageSCPD.xml
func SetSMBServer(session *soap.SoapSession, smbEnable bool) (tr064model.SetSMBServerResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_storage").
		Uri("urn:dslforum-org:service:X_AVM-DE_Storage:1").
		Action("SetSMBServer").
		AddBoolParam("NewSMBEnable", smbEnable).
		Do().Body.Data
	result := tr064model.SetSMBServerResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
