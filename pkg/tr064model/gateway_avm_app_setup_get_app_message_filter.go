package tr064model

import "encoding/xml"

// GetAppMessageFilterResponse AUTO-GENERATED (do not edit) model from [x_appsetupSCPD],
// based on SOAP action 'GetAppMessageFilter', Fritz!Box-System-Version 164.07.57
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
type GetAppMessageFilterResponse struct {
	XMLName    xml.Name // rather for debug purpose
	FilterList string   `xml:"NewFilterList"` //
}
