package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetDefaultConnectionService AUTO-GENERATED (do not edit) code from [layer3forwardingSCPD],
// based on SOAP action 'GetDefaultConnectionService', Fritz!Box-System-Version 164.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
func GetDefaultConnectionService(session *soap.SoapSession) (tr064model.GetDefaultConnectionServiceResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/layer3forwarding").
		Uri("urn:dslforum-org:service:Layer3Forwarding:1").
		Action("GetDefaultConnectionService").
		Do()
	if err != nil {
		return tr064model.GetDefaultConnectionServiceResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetDefaultConnectionServiceResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetDefaultConnectionServiceResponse{}, err
	}
	return result, nil
}
