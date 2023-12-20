package tr064model

import "encoding/xml"

// ConfigurationFinishedResponse AUTO-GENERATED (do not edit) model from [deviceconfigSCPD],
// based on SOAP action 'ConfigurationFinished', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
type ConfigurationFinishedResponse struct {
	XMLName xml.Name // rather for debug purpose
	Status  string   `xml:"NewStatus"` //
}
