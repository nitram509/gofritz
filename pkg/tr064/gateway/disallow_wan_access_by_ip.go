package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// DisallowWANAccessByIP AUTO-GENERATED (do not edit) code from [x_hostfilterSCPD],
// based on SOAP action 'DisallowWANAccessByIP', Fritz!Box-System-Version 164.07.57
//
// [x_hostfilterSCPD]: http://fritz.box:49000/x_hostfilterSCPD.xml
func DisallowWANAccessByIP(session *soap.SoapSession, ipv4Address string, disallow bool) (tr064model.DisallowWANAccessByIPResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_hostfilter").
		Uri("urn:dslforum-org:service:X_AVM-DE_HostFilter:1").
		Action("DisallowWANAccessByIP").
		AddStringParam("NewIPv4Address", ipv4Address).
		AddBoolParam("NewDisallow", disallow).
		Do()
	if err != nil {
		return tr064model.DisallowWANAccessByIPResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.DisallowWANAccessByIPResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.DisallowWANAccessByIPResponse{}, err
	}
	return result, nil
}
