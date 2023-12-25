package tr064model

import "encoding/xml"

// XavmGetDSLInfoResponse AUTO-GENERATED (do not edit) model from [wandslifconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetDSLInfo', Fritz!Box-System-Version 141.07.57
//
// [wandslifconfigSCPD]: http://fritz.box:49000/wandslifconfigSCPD.xml
type XavmGetDSLInfoResponse struct {
	XMLName               xml.Name // rather for debug purpose
	SNRGds                int      `xml:"NewSNRGds"`                // default=1
	SNRGus                int      `xml:"NewSNRGus"`                // default=1
	SNRpsds               string   `xml:"NewSNRpsds"`               //
	SNRpsus               string   `xml:"NewSNRpsus"`               //
	SNRMTds               int      `xml:"NewSNRMTds"`               // default=0
	SNRMTus               int      `xml:"NewSNRMTus"`               // default=0
	LATNds                string   `xml:"NewLATNds"`                //
	LATNus                string   `xml:"NewLATNus"`                //
	FECErrors             int      `xml:"NewFECErrors"`             // default=0
	CRCErrors             int      `xml:"NewCRCErrors"`             // default=0
	LinkStatus            string   `xml:"NewLinkStatus"`            // oneOf=[Up,Down,Initializing,Unavailable]
	ModulationType        string   `xml:"NewModulationType"`        // oneOf=[ADSL G.lite,G.shdsl,IDSL,HDSL,SDSL,VDSL]
	CurrentProfile        string   `xml:"NewCurrentProfile"`        //
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
