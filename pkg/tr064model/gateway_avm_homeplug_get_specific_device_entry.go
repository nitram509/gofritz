package tr064model

import "encoding/xml"

// GetSpecificDeviceEntryResponse AUTO-GENERATED (do not edit) model from [x_homeplugSCPD],
// based on SOAP action 'GetSpecificDeviceEntry', Fritz!Box-System-Version 164.08.00
//
// [x_homeplugSCPD]: http://fritz.box:49000/x_homeplugSCPD.xml
type GetSpecificDeviceEntryResponse struct {
	XMLName          xml.Name // rather for debug purpose
	Active           bool     `xml:"NewActive"`           // default=0
	Name             string   `xml:"NewName"`             //
	Model            string   `xml:"NewModel"`            //
	UpdateAvailable  bool     `xml:"NewUpdateAvailable"`  // default=0
	UpdateSuccessful string   `xml:"NewUpdateSuccessful"` // oneOf=[unknown,failed,succeeded]
}
