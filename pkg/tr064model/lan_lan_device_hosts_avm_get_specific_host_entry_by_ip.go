package tr064model

import "encoding/xml"

// XavmGetSpecificHostEntryByIPResponse AUTO-GENERATED (do not edit) model from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetSpecificHostEntryByIP', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
type XavmGetSpecificHostEntryByIPResponse struct {
	XMLName                 xml.Name // rather for debug purpose
	MACAddress              string   `xml:"NewMACAddress"`                       //
	Active                  bool     `xml:"NewActive"`                           // default=0
	HostName                string   `xml:"NewHostName"`                         //
	InterfaceType           string   `xml:"NewInterfaceType"`                    //
	Port                    int      `xml:"NewX_AVM-DE_Port"`                    // default=0
	Speed                   int      `xml:"NewX_AVM-DE_Speed"`                   // default=0
	UpdateAvailable         bool     `xml:"NewX_AVM-DE_UpdateAvailable"`         // default=0
	UpdateSuccessful        string   `xml:"NewX_AVM-DE_UpdateSuccessful"`        // default=unknown; oneOf=[unknown,failed,succeeded]
	InfoURL                 string   `xml:"NewX_AVM-DE_InfoURL"`                 //
	MACAddressList          string   `xml:"NewX_AVM-DE_MACAddressList"`          //
	Model                   string   `xml:"NewX_AVM-DE_Model"`                   //
	URL                     string   `xml:"NewX_AVM-DE_URL"`                     //
	Guest                   bool     `xml:"NewX_AVM-DE_Guest"`                   // default=0
	RequestClient           bool     `xml:"NewX_AVM-DE_RequestClient"`           // default=0
	VPN                     bool     `xml:"NewX_AVM-DE_VPN"`                     // default=0
	WANAccess               string   `xml:"NewX_AVM-DE_WANAccess"`               // default=unknown; oneOf=[unknown,error,granted,denied]
	Disallow                bool     `xml:"NewX_AVM-DE_Disallow"`                // default=0
	IsMeshable              bool     `xml:"NewX_AVM-DE_IsMeshable"`              // default=0
	Priority                bool     `xml:"NewX_AVM-DE_Priority"`                // default=0
	FriendlyName            string   `xml:"NewX_AVM-DE_FriendlyName"`            //
	FriendlyNameIsWriteable bool     `xml:"NewX_AVM-DE_FriendlyNameIsWriteable"` // default=0
}
