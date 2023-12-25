package tr064model

import "encoding/xml"

// GetAvmMediaInfoResponse AUTO-GENERATED (do not edit) model from [x_mediaSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 141.07.57
//
// [x_mediaSCPD]: http://fritz.box:49000/x_mediaSCPD.xml
type GetAvmMediaInfoResponse struct {
	XMLName             xml.Name // rather for debug purpose
	DVBCEnabled         bool     `xml:"NewDVBCEnabled"`         // default=0
	StationSearchStatus string   `xml:"NewStationSearchStatus"` // default=inactive; oneOf=[active,inactive]
	SearchProgress      int      `xml:"NewSearchProgress"`      // default=0
}
