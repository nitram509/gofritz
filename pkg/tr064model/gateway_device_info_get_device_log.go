package tr064model

import "encoding/xml"

// GetDeviceLogResponse AUTO-GENERATED (do not edit) model from [deviceinfoSCPD],
// based on SOAP action 'GetDeviceLog', Fritz!Box-System-Version 141.07.57
//
// [deviceinfoSCPD]: http://fritz.box:49000/deviceinfoSCPD.xml
type GetDeviceLogResponse struct {
	XMLName   xml.Name // rather for debug purpose
	DeviceLog string   `xml:"NewDeviceLog"` //
}
