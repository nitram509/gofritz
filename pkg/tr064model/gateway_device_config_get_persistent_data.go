package tr064model

import "encoding/xml"

// GetPersistentDataResponse AUTO-GENERATED (do not edit) model from [deviceconfigSCPD],
// based on SOAP action 'GetPersistentData', Fritz!Box-System-Version 141.07.57
//
// [deviceconfigSCPD]: http://fritz.box:49000/deviceconfigSCPD.xml
type GetPersistentDataResponse struct {
	XMLName        xml.Name // rather for debug purpose
	PersistentData string   `xml:"NewPersistentData"` //
}
