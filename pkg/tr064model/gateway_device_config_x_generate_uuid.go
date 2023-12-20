package tr064model

import "encoding/xml"

// X_GenerateUUIDResponse AUTO-GENERATED (do not edit) model from [deviceconfigSCPD],
// based on SOAP action 'X_GenerateUUID', Fritz!Box-System-Version 164.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
type X_GenerateUUIDResponse struct {
	XMLName xml.Name // rather for debug purpose
	UUID    string   `xml:"NewUUID"` //; type=UUID
}
