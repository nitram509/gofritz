package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetSubnetMask AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'SetSubnetMask', Fritz!Box-System-Version 164.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func SetSubnetMask(session *soap.SoapSession, subnetMask string) (tr064model.SetSubnetMaskResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("SetSubnetMask").
		AddStringParam("NewSubnetMask", subnetMask).
		Do()
	if err != nil {
		return tr064model.SetSubnetMaskResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetSubnetMaskResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetSubnetMaskResponse{}, err
	}
	return result, nil
}
