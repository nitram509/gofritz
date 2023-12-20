package tr064model

// GetAvmOnTelInfoResponse auto generated model from [x_contactSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetAvmOnTelInfoResponse struct {
	Enable      bool   `xml:"NewEnable"`      // default=0
	Status      string `xml:"NewStatus"`      //
	LastConnect string `xml:"NewLastConnect"` //
	Url         string `xml:"NewUrl"`         //
	ServiceId   string `xml:"NewServiceId"`   //
	Username    string `xml:"NewUsername"`    //
	Name        string `xml:"NewName"`        //
}