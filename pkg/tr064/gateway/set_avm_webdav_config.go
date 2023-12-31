package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetAvmWebdavConfig AUTO-GENERATED (do not edit) code from [x_webdavSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.07.57
//
// [x_webdavSCPD]: http://fritz.box:49000/x_webdavSCPD.xml
func SetAvmWebdavConfig(session *soap.SoapSession, enable bool, hostUrl string, username string, password string, mountpointName string) (tr064model.SetAvmWebdavConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_webdav").
		Uri("urn:dslforum-org:service:X_AVM-DE_WebDAVClient:1").
		Action("SetConfig").
		AddBoolParam("NewEnable", enable).
		AddStringParam("NewHostURL", hostUrl).
		AddStringParam("NewUsername", username).
		AddStringParam("NewPassword", password).
		AddStringParam("NewMountpointName", mountpointName).
		Do().Body.Data
	result := tr064model.SetAvmWebdavConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
