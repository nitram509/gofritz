package tr064model

import "encoding/xml"

// XavmGetDeviceLogPathResponse AUTO-GENERATED (do not edit) model from [deviceinfoSCPD],
// based on SOAP action 'X_AVM-DE_GetDeviceLogPath', Fritz!Box-System-Version 164.08.00
//
// [deviceinfoSCPD]: http://fritz.box:49000/deviceinfoSCPD.xml
type XavmGetDeviceLogPathResponse struct {
	XMLName       xml.Name // rather for debug purpose
	DeviceLogPath string   `xml:"NewDeviceLogPath"` //
}
