package tr064model

import "encoding/xml"

// SetAvmRemoteAccessEnableResponse AUTO-GENERATED (do not edit) model from [x_remoteSCPD],
// based on SOAP action 'SetEnable', Fritz!Box-System-Version 164.08.00
//
// [x_remoteSCPD]: http://fritz.box:49000/x_remoteSCPD.xml
type SetAvmRemoteAccessEnableResponse struct {
	XMLName xml.Name // rather for debug purpose
	Port    string   `xml:"NewPort"` //
}
