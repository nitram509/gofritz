package tr064model

import "encoding/xml"

// XavmGetNumberOfAlarmClocksResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetNumberOfAlarmClocks', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetNumberOfAlarmClocksResponse struct {
	XMLName             xml.Name // rather for debug purpose
	NumberOfAlarmClocks int      `xml:"NewX_AVM-DE_NumberOfAlarmClocks"` // default=3
}
