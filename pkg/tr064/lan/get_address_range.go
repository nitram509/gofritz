package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAddressRange AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetAddressRange', Fritz!Box-System-Version 164.08.00
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func GetAddressRange(session *soap.SoapSession) (tr064model.GetAddressRangeResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("GetAddressRange").
		Do()
	if err != nil {
		return tr064model.GetAddressRangeResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetAddressRangeResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetAddressRangeResponse{}, err
	}
	return result, nil
}
