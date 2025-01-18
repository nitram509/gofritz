package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// DiscardAllTickets AUTO-GENERATED (do not edit) code from [x_hostfilterSCPD],
// based on SOAP action 'DiscardAllTickets', Fritz!Box-System-Version 164.08.00
//
// [x_hostfilterSCPD]: http://fritz.box:49000/x_hostfilterSCPD.xml
func DiscardAllTickets(session *soap.SoapSession) (tr064model.DiscardAllTicketsResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_hostfilter").
		Uri("urn:dslforum-org:service:X_AVM-DE_HostFilter:1").
		Action("DiscardAllTickets").
		Do()
	if err != nil {
		return tr064model.DiscardAllTicketsResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.DiscardAllTicketsResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.DiscardAllTicketsResponse{}, err
	}
	return result, nil
}
