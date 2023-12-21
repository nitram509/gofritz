package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetAvmRemoteAccessInfo AUTO-GENERATED (do not edit) code from [x_remoteSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
func GetAvmRemoteAccessInfo(session *soap.SoapSession) (tr064model.GetAvmRemoteAccessInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_remote").
		Uri("urn:dslforum-org:service:X_AVM-DE_RemoteAccess:1").
		Action("GetInfo").
		Do().Body.Data
	result := tr064model.GetAvmRemoteAccessInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
