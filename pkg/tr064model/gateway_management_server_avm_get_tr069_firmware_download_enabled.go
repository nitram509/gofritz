package tr064model

import "encoding/xml"

// XavmGetTR069FirmwareDownloadEnabledResponse AUTO-GENERATED (do not edit) model from [mgmsrvSCPD],
// based on SOAP action 'X_AVM-DE_GetTR069FirmwareDownloadEnabled', Fritz!Box-System-Version 164.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
type XavmGetTR069FirmwareDownloadEnabledResponse struct {
	XMLName                      xml.Name // rather for debug purpose
	TR069FirmwareDownloadEnabled bool     `xml:"NewTR069FirmwareDownloadEnabled"` // default=0
}
