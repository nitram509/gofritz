package tr064model

import "encoding/xml"

// XavmGetAlarmClockResponse AUTO-GENERATED (do not edit) model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetAlarmClock', Fritz!Box-System-Version 164.08.00
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetAlarmClockResponse struct {
	XMLName             xml.Name // rather for debug purpose
	AlarmClockEnable    bool     `xml:"NewX_AVM-DE_AlarmClockEnable"`    // default=0
	AlarmClockName      string   `xml:"NewX_AVM-DE_AlarmClockName"`      //
	AlarmClockTime      string   `xml:"NewX_AVM-DE_AlarmClockTime"`      //
	AlarmClockWeekdays  string   `xml:"NewX_AVM-DE_AlarmClockWeekdays"`  //
	AlarmClockPhoneName string   `xml:"NewX_AVM-DE_AlarmClockPhoneName"` //
}
