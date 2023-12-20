package tr064model

import "encoding/xml"

// GetWanDslIfConfigInfoResponse AUTO-GENERATED (do not edit) model from [wandslifconfigSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [wandslifconfigSCPD]: http://fritz.box:49000/wandslifconfigSCPD.xml
type GetWanDslIfConfigInfoResponse struct {
	XMLName               xml.Name // rather for debug purpose
	Enable                bool     `xml:"NewEnable"`                // default=0
	Status                string   `xml:"NewStatus"`                //
	DataPath              string   `xml:"NewDataPath"`              //
	UpstreamCurrRate      int      `xml:"NewUpstreamCurrRate"`      // default=0
	DownstreamCurrRate    int      `xml:"NewDownstreamCurrRate"`    // default=0
	UpstreamMaxRate       int      `xml:"NewUpstreamMaxRate"`       // default=0
	DownstreamMaxRate     int      `xml:"NewDownstreamMaxRate"`     // default=0
	UpstreamNoiseMargin   int      `xml:"NewUpstreamNoiseMargin"`   // default=0
	DownstreamNoiseMargin int      `xml:"NewDownstreamNoiseMargin"` // default=0
	UpstreamAttenuation   int      `xml:"NewUpstreamAttenuation"`   // default=0
	DownstreamAttenuation int      `xml:"NewDownstreamAttenuation"` // default=0
	ATURVendor            string   `xml:"NewATURVendor"`            //
	ATURCountry           string   `xml:"NewATURCountry"`           //
	UpstreamPower         int      `xml:"NewUpstreamPower"`         // default=0
	DownstreamPower       int      `xml:"NewDownstreamPower"`       // default=0
}
