package tr064model

import "encoding/xml"

// XavmGetSupportDataInfoResponse AUTO-GENERATED (do not edit) model from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetSupportDataInfo', Fritz!Box-System-Version 164.08.00
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
type XavmGetSupportDataInfoResponse struct {
	XMLName              xml.Name // rather for debug purpose
	SupportDataMode      string   `xml:"NewX_AVM-DE_SupportDataMode"`      // oneOf=[normal,mesh,unknown]
	SupportDataID        string   `xml:"NewX_AVM-DE_SupportDataID"`        //
	SupportDataTimestamp string   `xml:"NewX_AVM-DE_SupportDataTimestamp"` // default=0001-01-01T00:00:00; type=Datetime
	SupportDataStatus    string   `xml:"NewX_AVM-DE_SupportDataStatus"`    // oneOf=[unknown,blocked,ok,preparing,error,creating]
	SupportDataEnabled   bool     `xml:"NewX_AVM-DE_SupportDataEnabled"`   // default=0
}
