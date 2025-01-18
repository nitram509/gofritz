package tr064model

import "encoding/xml"

// GetPersistentDataResponse AUTO-GENERATED (do not edit) model from [deviceconfigSCPD],
// based on SOAP action 'GetPersistentData', Fritz!Box-System-Version 164.08.00
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
type GetPersistentDataResponse struct {
	XMLName        xml.Name // rather for debug purpose
	PersistentData string   `xml:"NewPersistentData"` //
}
