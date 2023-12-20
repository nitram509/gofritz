package tr064model

// RequestFTPServerWANResponse auto generated model from [x_storageSCPD],
// based on SOAP action 'RequestFTPServerWAN', Fritz!Box-System-Version 164.07.57
//
// [x_storageSCPD]: http://fritz.box:49000/x_storageSCPD.xml
type RequestFTPServerWANResponse struct {
	FTPWANPort     int `xml:"NewFTPWANPort"`     // default=0
	FTPWANLifetime int `xml:"NewFTPWANLifetime"` // default=0
}
