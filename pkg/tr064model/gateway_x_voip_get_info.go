package tr064model

// GetXVoipInfoResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type GetXVoipInfoResponse struct {
	FaxT38Enable bool   `xml:"NewFaxT38Enable"` // default=0
	VoiceCoding  string `xml:"NewVoiceCoding"`  // oneOf=[auto,fixed,compressed,autocompressed]
}
