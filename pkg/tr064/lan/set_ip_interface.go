package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetIPInterface AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'SetIPInterface', Fritz!Box-System-Version 164.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func SetIPInterface(session *soap.SoapSession, enable bool, ipAddress string, subnetMask string, ipAddressingType string) (tr064model.SetIPInterfaceResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("SetIPInterface").
		AddBoolParam("NewEnable", enable).
		AddStringParam("NewIPAddress", ipAddress).
		AddStringParam("NewSubnetMask", subnetMask).
		AddStringParam("NewIPAddressingType", ipAddressingType).
		Do()
	if err != nil {
		return tr064model.SetIPInterfaceResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetIPInterfaceResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetIPInterfaceResponse{}, err
	}
	return result, nil
}
