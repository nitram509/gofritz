package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetDHCPServerEnable AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'SetDHCPServerEnable', Fritz!Box-System-Version 164.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func SetDHCPServerEnable(session *soap.SoapSession, dhcpServerEnable bool) (tr064model.SetDHCPServerEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("SetDHCPServerEnable").
		AddBoolParam("NewDHCPServerEnable", dhcpServerEnable).
		Do().Body.Data
	result := tr064model.SetDHCPServerEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
