package tr064model

import "encoding/xml"

// GetTimeInfoResponse AUTO-GENERATED (do not edit) model from [timeSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.08.00
//
// [timeSCPD]: http://fritz.box:49000/timeSCPD.xml
type GetTimeInfoResponse struct {
	XMLName              xml.Name // rather for debug purpose
	NTPServer1           string   `xml:"NewNTPServer1"`           //
	NTPServer2           string   `xml:"NewNTPServer2"`           //
	CurrentLocalTime     string   `xml:"NewCurrentLocalTime"`     // default=0001-01-01T00:00:00; type=Datetime
	LocalTimeZone        string   `xml:"NewLocalTimeZone"`        //
	LocalTimeZoneName    string   `xml:"NewLocalTimeZoneName"`    //
	DaylightSavingsUsed  bool     `xml:"NewDaylightSavingsUsed"`  // default=0
	DaylightSavingsStart string   `xml:"NewDaylightSavingsStart"` // default=0001-01-01T00:00:00; type=Datetime
	DaylightSavingsEnd   string   `xml:"NewDaylightSavingsEnd"`   // default=0001-01-01T00:00:00; type=Datetime
}
