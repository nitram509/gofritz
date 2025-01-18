package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetAvmWebdavInfo AUTO-GENERATED (do not edit) code from [x_webdavSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [x_webdavSCPD]: http://fritz.box:49000/x_webdavSCPD.xml
func GetAvmWebdavInfo(session *soap.SoapSession) (tr064model.GetAvmWebdavInfoResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_webdav").
		Uri("urn:dslforum-org:service:X_AVM-DE_WebDAVClient:1").
		Action("GetInfo").
		Do()
	if err != nil {
		return tr064model.GetAvmWebdavInfoResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetAvmWebdavInfoResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetAvmWebdavInfoResponse{}, err
	}
	return result, nil
}
