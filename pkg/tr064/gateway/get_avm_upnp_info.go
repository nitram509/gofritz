package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAvmUpnpInfo AUTO-GENERATED (do not edit) code from [x_upnpSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_upnpSCPD]: http://fritz.box:49000/x_upnpSCPD.xml
func GetAvmUpnpInfo(session *soap.SoapSession) (tr064model.GetAvmUpnpInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_upnp").
		Uri("urn:dslforum-org:service:X_AVM-DE_UPnP:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetAvmUpnpInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetAvmUpnpInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetAvmUpnpInfoResponse{}, err
	}
	return result, nil
}
