package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetSMBServer AUTO-GENERATED (do not edit) code from [x_storageSCPD],
// based on SOAP action 'SetSMBServer', Fritz!Box-System-Version 164.08.00
//
// [x_storageSCPD]: http://fritz.box:49000/x_storageSCPD.xml
func SetSMBServer(session *soap.SoapSession, smbEnable bool) (tr064model.SetSMBServerResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_storage").
		Uri("urn:dslforum-org:service:X_AVM-DE_Storage:1").
		Action("SetSMBServer").
		AddBoolParam("NewSMBEnable", smbEnable).
		Do()
	if err != nil {
		return tr064model.SetSMBServerResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetSMBServerResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetSMBServerResponse{}, err
	}
	return result, nil
}
