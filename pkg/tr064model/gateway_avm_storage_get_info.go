package tr064model

import "encoding/xml"

// GetAvmStorageInfoResponse AUTO-GENERATED (do not edit) model from [x_storageSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [x_storageSCPD]: http://fritz.box:49000/x_storageSCPD.xml
type GetAvmStorageInfoResponse struct {
	XMLName       xml.Name // rather for debug purpose
	FTPEnable     bool     `xml:"NewFTPEnable"`     // default=0
	FTPStatus     string   `xml:"NewFTPStatus"`     //
	SMBEnable     bool     `xml:"NewSMBEnable"`     // default=0
	FTPWANEnable  bool     `xml:"NewFTPWANEnable"`  // default=0
	FTPWANSSLOnly bool     `xml:"NewFTPWANSSLOnly"` // default=0
	FTPWANPort    int      `xml:"NewFTPWANPort"`    // default=0
}
