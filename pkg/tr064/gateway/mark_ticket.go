package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// MarkTicket AUTO-GENERATED (do not edit) code from [x_hostfilterSCPD],
// based on SOAP action 'MarkTicket', Fritz!Box-System-Version 164.07.57
//
// [x_hostfilterSCPD]: http://fritz.box:49000/x_hostfilterSCPD.xml
func MarkTicket(session *soap.SoapSession) (tr064model.MarkTicketResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_hostfilter").
		Uri("urn:dslforum-org:service:X_AVM-DE_HostFilter:1").
		Action("MarkTicket").
		Do()
	if err != nil {
		return tr064model.MarkTicketResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.MarkTicketResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.MarkTicketResponse{}, err
	}
	return result, nil
}
