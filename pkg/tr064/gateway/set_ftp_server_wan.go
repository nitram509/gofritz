package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetFTPServerWAN AUTO-GENERATED (do not edit) code from [x_storageSCPD],
// based on SOAP action 'SetFTPServerWAN', Fritz!Box-System-Version 164.08.00
//
// [x_storageSCPD]: http://fritz.box:49000/x_storageSCPD.xml
func SetFTPServerWAN(session *soap.SoapSession, ftpWanEnable bool, ftpWanSSLOnly bool) (tr064model.SetFTPServerWANResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_storage").
		Uri("urn:dslforum-org:service:X_AVM-DE_Storage:1").
		Action("SetFTPServerWAN").
		AddBoolParam("NewFTPWANEnable", ftpWanEnable).
		AddBoolParam("NewFTPWANSSLOnly", ftpWanSSLOnly).
		Do()
	if err != nil {
		return tr064model.SetFTPServerWANResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetFTPServerWANResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetFTPServerWANResponse{}, err
	}
	return result, nil
}
