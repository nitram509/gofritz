package tr064model

import "encoding/xml"

// GetXVoipInfoResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type GetXVoipInfoResponse struct {
	XMLName      xml.Name // rather for debug purpose
	FaxT38Enable bool     `xml:"NewFaxT38Enable"` // default=0
	VoiceCoding  string   `xml:"NewVoiceCoding"`  // oneOf=[auto,fixed,compressed,autocompressed]
}
