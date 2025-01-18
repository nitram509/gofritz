package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetDDNSInfo AUTO-GENERATED (do not edit) code from [x_remoteSCPD],
// based on SOAP action 'GetDDNSInfo', Fritz!Box-System-Version 164.08.00
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
func GetDDNSInfo(session *soap.SoapSession) (tr064model.GetDDNSInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_remote").
		Uri("urn:dslforum-org:service:X_AVM-DE_RemoteAccess:1").
		Action("GetDDNSInfo").
		Do()
	if err != nil {
		return tr064model.GetDDNSInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetDDNSInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetDDNSInfoResponse{}, err
	}
	return result, nil
}
