package tr064model

import "encoding/xml"

// AddUSPControllerResponse AUTO-GENERATED (do not edit) model from [x_uspcontrollerSCPD],
// based on SOAP action 'AddUSPController', Fritz!Box-System-Version 141.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
type AddUSPControllerResponse struct {
	XMLName xml.Name // rather for debug purpose
	Index   int      `xml:"NewIndex"` // default=0
}
