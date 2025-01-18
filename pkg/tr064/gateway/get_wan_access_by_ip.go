package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetWANAccessByIP AUTO-GENERATED (do not edit) code from [x_hostfilterSCPD],
// based on SOAP action 'GetWANAccessByIP', Fritz!Box-System-Version 164.08.00
//
// [x_hostfilterSCPD]: http://fritz.box:49000/x_hostfilterSCPD.xml
func GetWANAccessByIP(session *soap.SoapSession, ipv4Address string) (tr064model.GetWANAccessByIPResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_hostfilter").
		Uri("urn:dslforum-org:service:X_AVM-DE_HostFilter:1").
		Action("GetWANAccessByIP").
		AddStringParam("NewIPv4Address", ipv4Address).
		Do()
	if err != nil {
		return tr064model.GetWANAccessByIPResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetWANAccessByIPResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetWANAccessByIPResponse{}, err
	}
	return result, nil
}
