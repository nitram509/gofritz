package tr064model

// GetAvmMyFritzInfoResponse AUTO-GENERATED (do not edit) model from [x_myfritzSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_myfritzSCPD]: http://fritz.box:49000/x_myfritzSCPD.xml
type GetAvmMyFritzInfoResponse struct {
	Enabled          bool   `xml:"NewEnabled"`          // default=0
	DeviceRegistered bool   `xml:"NewDeviceRegistered"` // default=0
	DynDNSName       string `xml:"NewDynDNSName"`       //
	Port             int    `xml:"NewPort"`             // default=0
	State            string `xml:"NewState"`            //
	Email            string `xml:"NewEmail"`            //
}
