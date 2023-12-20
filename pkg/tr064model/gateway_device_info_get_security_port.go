package tr064model

import "encoding/xml"

// GetSecurityPortResponse AUTO-GENERATED (do not edit) model from [deviceinfoSCPD],
// based on SOAP action 'GetSecurityPort', Fritz!Box-System-Version 164.07.57
//
// [deviceinfoSCPD]: http://fritz.box:49000/deviceinfoSCPD.xml
type GetSecurityPortResponse struct {
	XMLName      xml.Name // rather for debug purpose
	SecurityPort int      `xml:"NewSecurityPort"` // default=0
}
