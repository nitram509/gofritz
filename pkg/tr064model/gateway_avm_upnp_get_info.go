package tr064model

// GetAvmUpnpInfoResponse AUTO-GENERATED (do not edit) model from [x_upnpSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_upnpSCPD]: http://fritz.box:49000/x_upnpSCPD.xml
type GetAvmUpnpInfoResponse struct {
	Enable          bool `xml:"NewEnable"`          // default=0
	UPnPMediaServer bool `xml:"NewUPnPMediaServer"` // default=0
}
