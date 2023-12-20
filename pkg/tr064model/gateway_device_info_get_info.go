package tr064model

import "encoding/xml"

// GetDeviceInfoResponse AUTO-GENERATED (do not edit) model from [deviceinfoSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [deviceinfoSCPD]: http://fritz.box:49000/deviceinfoSCPD.xml
type GetDeviceInfoResponse struct {
	XMLName          xml.Name // rather for debug purpose
	ManufacturerName string   `xml:"NewManufacturerName"` // default=AVM
	ManufacturerOUI  string   `xml:"NewManufacturerOUI"`  //
	ModelName        string   `xml:"NewModelName"`        //
	Description      string   `xml:"NewDescription"`      //
	ProductClass     string   `xml:"NewProductClass"`     //
	SerialNumber     string   `xml:"NewSerialNumber"`     //
	SoftwareVersion  string   `xml:"NewSoftwareVersion"`  //
	HardwareVersion  string   `xml:"NewHardwareVersion"`  //
	SpecVersion      string   `xml:"NewSpecVersion"`      //
	ProvisioningCode string   `xml:"NewProvisioningCode"` //
	UpTime           int      `xml:"NewUpTime"`           // default=0
	DeviceLog        string   `xml:"NewDeviceLog"`        //
}
