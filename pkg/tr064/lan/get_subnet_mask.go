package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetSubnetMask AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetSubnetMask', Fritz!Box-System-Version 141.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func GetSubnetMask(session *soap.SoapSession) (tr064model.GetSubnetMaskResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("GetSubnetMask").
		Do().Body.Data
	result := tr064model.GetSubnetMaskResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
