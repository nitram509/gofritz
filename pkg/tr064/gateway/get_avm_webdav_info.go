package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetAvmWebdavInfo AUTO-GENERATED (do not edit) code from [x_webdavSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 141.07.57
//
// [x_webdavSCPD]: http://fritz.box:49000/x_webdavSCPD.xml
func GetAvmWebdavInfo(session *soap.SoapSession) (tr064model.GetAvmWebdavInfoResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_webdav").
		Uri("urn:dslforum-org:service:X_AVM-DE_WebDAVClient:1").
		Action("GetInfo").
		Do().Body.Data
	result := tr064model.GetAvmWebdavInfoResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
