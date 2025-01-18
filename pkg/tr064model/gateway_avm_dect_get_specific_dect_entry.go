package tr064model

import "encoding/xml"

// GetSpecificDectEntryResponse AUTO-GENERATED (do not edit) model from [x_dectSCPD],
// based on SOAP action 'GetSpecificDectEntry', Fritz!Box-System-Version 164.08.00
//
// [x_dectSCPD]: http://fritz.box:49000/x_dectSCPD.xml
type GetSpecificDectEntryResponse struct {
	XMLName          xml.Name // rather for debug purpose
	Active           bool     `xml:"NewActive"`           // default=0
	Name             string   `xml:"NewName"`             //
	Model            string   `xml:"NewModel"`            //
	UpdateAvailable  bool     `xml:"NewUpdateAvailable"`  // default=0
	UpdateSuccessful string   `xml:"NewUpdateSuccessful"` // oneOf=[unknown,failed,succeeded]
	UpdateInfo       string   `xml:"NewUpdateInfo"`       //
}
