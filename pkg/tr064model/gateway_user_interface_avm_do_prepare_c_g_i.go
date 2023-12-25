package tr064model

import "encoding/xml"

// XavmDoPrepareCGIResponse AUTO-GENERATED (do not edit) model from [userifSCPD],
// based on SOAP action 'X_AVM-DE_DoPrepareCGI', Fritz!Box-System-Version 141.07.57
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
type XavmDoPrepareCGIResponse struct {
	XMLName   xml.Name // rather for debug purpose
	CGI       string   `xml:"NewX_AVM-DE_CGI"`       //
	SessionID string   `xml:"NewX_AVM-DE_SessionID"` //
}
