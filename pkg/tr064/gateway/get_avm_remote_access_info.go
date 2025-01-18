package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAvmRemoteAccessInfo AUTO-GENERATED (do not edit) code from [x_remoteSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
func GetAvmRemoteAccessInfo(session *soap.SoapSession) (tr064model.GetAvmRemoteAccessInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_remote").
		Uri("urn:dslforum-org:service:X_AVM-DE_RemoteAccess:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetAvmRemoteAccessInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetAvmRemoteAccessInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetAvmRemoteAccessInfoResponse{}, err
	}
	return result, nil
}
