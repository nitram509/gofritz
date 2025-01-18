package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetSubnetMask AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetSubnetMask', Fritz!Box-System-Version 164.08.00
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func GetSubnetMask(session *soap.SoapSession) (tr064model.GetSubnetMaskResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("GetSubnetMask").
		Do()
	if err != nil {
		return tr064model.GetSubnetMaskResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetSubnetMaskResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetSubnetMaskResponse{}, err
	}
	return result, nil
}
