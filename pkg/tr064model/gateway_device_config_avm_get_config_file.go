package tr064model

import "encoding/xml"

// XavmGetConfigFileResponse AUTO-GENERATED (do not edit) model from [deviceconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetConfigFile', Fritz!Box-System-Version 164.08.00
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
type XavmGetConfigFileResponse struct {
	XMLName       xml.Name // rather for debug purpose
	ConfigFileUrl string   `xml:"NewX_AVM-DE_ConfigFileUrl"` //
}
