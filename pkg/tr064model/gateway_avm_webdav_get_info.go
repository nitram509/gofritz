package tr064model

import "encoding/xml"

// GetAvmWebdavInfoResponse AUTO-GENERATED (do not edit) model from [x_webdavSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 141.07.57
//
// [x_webdavSCPD]: http://fritz.box:49000/x_webdavSCPD.xml
type GetAvmWebdavInfoResponse struct {
	XMLName        xml.Name // rather for debug purpose
	Enable         bool     `xml:"NewEnable"`         // default=0
	HostURL        string   `xml:"NewHostURL"`        //
	Username       string   `xml:"NewUsername"`       //
	MountpointName string   `xml:"NewMountpointName"` //
}
