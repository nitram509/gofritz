package tr064model

// XavmGetAlarmClockResponse auto generated model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetAlarmClock', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetAlarmClockResponse struct {
	AlarmClockEnable    bool   `xml:"NewX_AVM-DE_AlarmClockEnable"`    // default=0
	AlarmClockName      string `xml:"NewX_AVM-DE_AlarmClockName"`      //
	AlarmClockTime      string `xml:"NewX_AVM-DE_AlarmClockTime"`      //
	AlarmClockWeekdays  string `xml:"NewX_AVM-DE_AlarmClockWeekdays"`  //
	AlarmClockPhoneName string `xml:"NewX_AVM-DE_AlarmClockPhoneName"` //
}
