package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetAddressRange AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetAddressRange', Fritz!Box-System-Version 164.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func GetAddressRange(session *soap.SoapSession) (tr064model.GetAddressRangeResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("GetAddressRange").
		Do().Body.Data
	result := tr064model.GetAddressRangeResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
