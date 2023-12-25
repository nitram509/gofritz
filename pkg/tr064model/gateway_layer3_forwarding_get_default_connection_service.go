package tr064model

import "encoding/xml"

// GetDefaultConnectionServiceResponse AUTO-GENERATED (do not edit) model from [layer3forwardingSCPD],
// based on SOAP action 'GetDefaultConnectionService', Fritz!Box-System-Version 141.07.57
//
// [layer3forwardingSCPD]: http://fritz.box:49000/layer3forwardingSCPD.xml
type GetDefaultConnectionServiceResponse struct {
	XMLName                  xml.Name // rather for debug purpose
	DefaultConnectionService string   `xml:"NewDefaultConnectionService"` //
}
