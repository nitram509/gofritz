package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetAvmUpnpConfig AUTO-GENERATED (do not edit) code from [x_upnpSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.08.00
//
// [x_upnpSCPD]: http://fritz.box:49000/x_upnpSCPD.xml
func SetAvmUpnpConfig(session *soap.SoapSession, enable bool, upnpMediaServer bool) (tr064model.SetAvmUpnpConfigResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_upnp").
		Uri("urn:dslforum-org:service:X_AVM-DE_UPnP:1").
		Action("SetConfig").
		AddBoolParam("NewEnable", enable).
		AddBoolParam("NewUPnPMediaServer", upnpMediaServer).
		Do()
	if err != nil {
		return tr064model.SetAvmUpnpConfigResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetAvmUpnpConfigResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetAvmUpnpConfigResponse{}, err
	}
	return result, nil
}
