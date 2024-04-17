package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetDHCPServerEnable AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'SetDHCPServerEnable', Fritz!Box-System-Version 164.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func SetDHCPServerEnable(session *soap.SoapSession, dhcpServerEnable bool) (tr064model.SetDHCPServerEnableResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("SetDHCPServerEnable").
		AddBoolParam("NewDHCPServerEnable", dhcpServerEnable).
		Do()
	if err != nil {
		return tr064model.SetDHCPServerEnableResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetDHCPServerEnableResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetDHCPServerEnableResponse{}, err
	}
	return result, nil
}
