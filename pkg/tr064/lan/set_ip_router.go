package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetIPRouter AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'SetIPRouter', Fritz!Box-System-Version 164.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func SetIPRouter(session *soap.SoapSession, ipRouters string) (tr064model.SetIPRouterResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("SetIPRouter").
		AddStringParam("NewIPRouters", ipRouters).
		Do()
	if err != nil {
		return tr064model.SetIPRouterResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetIPRouterResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetIPRouterResponse{}, err
	}
	return result, nil
}
