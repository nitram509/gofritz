package tr064model

// GetAvmWebdavInfoResponse auto generated model from [x_webdavSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_webdavSCPD]: http://fritz.box:49000/x_webdavSCPD.xml
type GetAvmWebdavInfoResponse struct {
	Enable         bool   `xml:"NewEnable"`         // default=0
	HostURL        string `xml:"NewHostURL"`        //
	Username       string `xml:"NewUsername"`       //
	MountpointName string `xml:"NewMountpointName"` //
}
