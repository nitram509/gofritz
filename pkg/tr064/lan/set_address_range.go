package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetAddressRange AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'SetAddressRange', Fritz!Box-System-Version 164.08.00
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func SetAddressRange(session *soap.SoapSession, minAddress string, maxAddress string) (tr064model.SetAddressRangeResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("SetAddressRange").
		AddStringParam("NewMinAddress", minAddress).
		AddStringParam("NewMaxAddress", maxAddress).
		Do()
	if err != nil {
		return tr064model.SetAddressRangeResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetAddressRangeResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetAddressRangeResponse{}, err
	}
	return result, nil
}
