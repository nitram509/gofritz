package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetProvisioningCode AUTO-GENERATED (do not edit) code from [deviceinfoSCPD],
// based on SOAP action 'SetProvisioningCode', Fritz!Box-System-Version 141.07.57
//
// [deviceinfoSCPD]: http://fritz.box:49000/deviceinfoSCPD.xml
func SetProvisioningCode(session *soap.SoapSession, provisioningCode string) (tr064model.SetProvisioningCodeResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/deviceinfo").
		Uri("urn:dslforum-org:service:DeviceInfo:1").
		Action("SetProvisioningCode").
		AddStringParam("NewProvisioningCode", provisioningCode).
		Do().Body.Data
	result := tr064model.SetProvisioningCodeResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
