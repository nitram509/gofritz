package tr064model

import "encoding/xml"

// RequestFTPServerWANResponse AUTO-GENERATED (do not edit) model from [x_storageSCPD],
// based on SOAP action 'RequestFTPServerWAN', Fritz!Box-System-Version 164.08.00
//
// [x_storageSCPD]: http://fritz.box:49000/x_storageSCPD.xml
type RequestFTPServerWANResponse struct {
	XMLName        xml.Name // rather for debug purpose
	FTPWANPort     int      `xml:"NewFTPWANPort"`     // default=0
	FTPWANLifetime int      `xml:"NewFTPWANLifetime"` // default=0
}
