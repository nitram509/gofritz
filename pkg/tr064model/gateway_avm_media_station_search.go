package tr064model

import "encoding/xml"

// StationSearchResponse AUTO-GENERATED (do not edit) model from [x_mediaSCPD],
// based on SOAP action 'StationSearch', Fritz!Box-System-Version 141.07.57
//
// [x_mediaSCPD]: http://fritz.box:49000/x_mediaSCPD.xml
type StationSearchResponse struct {
	XMLName             xml.Name // rather for debug purpose
	StationSearchStatus string   `xml:"NewStationSearchStatus"` // default=inactive; oneOf=[active,inactive]
}
