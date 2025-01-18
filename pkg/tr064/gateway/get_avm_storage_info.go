package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAvmStorageInfo AUTO-GENERATED (do not edit) code from [x_storageSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [x_storageSCPD]: http://fritz.box:49000/x_storageSCPD.xml
func GetAvmStorageInfo(session *soap.SoapSession) (tr064model.GetAvmStorageInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_storage").
		Uri("urn:dslforum-org:service:X_AVM-DE_Storage:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetAvmStorageInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetAvmStorageInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetAvmStorageInfoResponse{}, err
	}
	return result, nil
}
