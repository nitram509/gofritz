package tr064model

import "encoding/xml"

// SetAvmAuthConfigResponse AUTO-GENERATED (do not edit) model from [x_authSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.08.00
//
// [x_authSCPD]: http://fritz.box:49000/x_authSCPD.xml
type SetAvmAuthConfigResponse struct {
	XMLName xml.Name // rather for debug purpose
	State   string   `xml:"NewState"`   //
	Token   string   `xml:"NewToken"`   //
	Methods string   `xml:"NewMethods"` //
}
