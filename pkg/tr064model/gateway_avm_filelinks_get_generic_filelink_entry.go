package tr064model

// GetGenericFilelinkEntryResponse auto generated model from [x_filelinksSCPD],
// based on SOAP action 'GetGenericFilelinkEntry', Fritz!Box-System-Version 164.07.57
//
// [x_filelinksSCPD]: http://fritz.box:49000/x_filelinksSCPD.xml
type GetGenericFilelinkEntryResponse struct {
	ID               string `xml:"NewID"`               //
	Valid            bool   `xml:"NewValid"`            // default=0
	Path             string `xml:"NewPath"`             //
	IsDirectory      bool   `xml:"NewIsDirectory"`      // default=0
	Url              string `xml:"NewUrl"`              //
	Username         string `xml:"NewUsername"`         //
	AccessCountLimit int    `xml:"NewAccessCountLimit"` // default=0
	AccessCount      int    `xml:"NewAccessCount"`      // default=0
	Expire           int    `xml:"NewExpire"`           // default=0
	ExpireDate       string `xml:"NewExpireDate"`       // default=0001-01-01T00:00:00; type=Datetime
}
