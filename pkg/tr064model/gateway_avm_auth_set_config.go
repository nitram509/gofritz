package tr064model

// SetAvmAuthConfigResponse auto generated model from [x_authSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.07.57
//
// [x_authSCPD]: http://fritz.box:49000/x_authSCPD.xml
type SetAvmAuthConfigResponse struct {
	State   string `xml:"NewState"`   //
	Token   string `xml:"NewToken"`   //
	Methods string `xml:"NewMethods"` //
}
