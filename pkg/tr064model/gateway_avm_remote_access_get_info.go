package tr064model

// GetAvmRemoteAccessInfoResponse auto generated model from [x_remoteSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
type GetAvmRemoteAccessInfoResponse struct {
	Enabled            bool   `xml:"NewEnabled"`            // default=0
	Port               string `xml:"NewPort"`               //
	Username           string `xml:"NewUsername"`           //
	LetsEncryptEnabled bool   `xml:"NewLetsEncryptEnabled"` // default=0
	LetsEncryptState   string `xml:"NewLetsEncryptState"`   //
}
